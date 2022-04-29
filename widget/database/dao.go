package database

import (
	"database/sql"
	"fmt"
	"github.com/whf-sky/efficient/widget/dtype"
	"github.com/whf-sky/efficient/widget/validation"
	"log"
)

type ModelHandle func() ModelInterface

type Dao struct {
	db          *sql.DB
	driverName  string
	tx          *sql.Tx
	data        []ModelInterface
	modelHandle func() ModelInterface
	model       ModelInterface
	where       string
	whereArgs   []interface{}
	rows        *sql.Rows
	row         *sql.Row
	error       error
}

func (d *Dao) SetDb(db *sql.DB) *sql.DB {
	d.db = db
	return d.db
}

type DaoInterface interface {
}

func (d *Dao) DriverName(driverName string) {
	d.driverName = driverName
}

func (d *Dao) SetModel(fn func() ModelInterface) *Dao {
	d.modelHandle = fn
	d.model = fn()
	return d
}

func (d *Dao) Model() ModelInterface {
	return d.model
}

func (d *Dao) SetData(data ...ModelInterface) *Dao {
	d.data = data
	return d
}

func (d *Dao) Where(where string, args ...interface{}) *Dao {
	d.where = where
	d.whereArgs = args
	return d
}

func (d *Dao) Insert() (int64, error) {
	columns, values, err := d.values()
	if err != nil {
		return 0, err
	}
	vLen := len(values)
	sql, err := SQL(d.driverName).InsertQuery(d.TableName(), columns, vLen)
	if err != nil {
		return 0, err
	}
	result, err := d.Exec(sql, values...)
	if err != nil {
		return 0, err
	}
	if len(columns) == vLen {
		return result.LastInsertId()
	}
	return result.RowsAffected()
}

func (d *Dao) TableName(indexs ...int) string {
	if len(d.data) > 0 {
		index := 0
		if len(indexs) >0 {
			index = indexs[0]
		}
		return d.data[index].TableName()
	}
	return d.model.TableName()
}

func (d *Dao) Update() (rowsAffected int64, err error) {
	columns, values, err := d.sets()
	sql, err := SQL(d.driverName).UpdateQuery(d.TableName(), columns, d.where)
	if err != nil {
		return 0, err
	}
	args := append(values, d.whereArgs...)
	fmt.Println(sql, args)
	result, err := d.Exec(sql, args...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (d *Dao) Delete() (rowsAffected int64, err error) {
	sql, err := SQL(d.driverName).DeleteQuery(d.TableName(), d.where)
	if err != nil {
		return 0, err
	}
	fmt.Println(sql, d.whereArgs)
	result, err := d.Exec(sql, d.whereArgs...)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (d *Dao) QueryRow(query string, args ...interface{}) *Dao {
	if d.tx != nil {
		d.row = d.tx.QueryRow(query, args...)
	} else {
		d.row = d.db.QueryRow(query, args...)
	}
	return d
}

func (d *Dao) Row(dest ...interface{}) error {
	err := d.row.Scan(dest...)
	return err
}

func (d *Dao) Query(query string, args ...interface{}) *Dao {
	if d.tx == nil {
		d.rows, d.error = d.db.Query(query, args...)
	} else {
		d.rows, d.error = d.tx.Query(query, args...)
	}
	return d
}

func (d *Dao) Rows() (*sql.Rows, error) {
	return d.rows, d.error
}

func (d *Dao) ToModel() (model ModelInterface, err error) {
	models, err := d.ToModels()
	if len(models) == 0 {
		return nil, err
	}
	return models[0], err
}

func (d *Dao) ToModels() (models []ModelInterface, err error) {
	defer d.rows.Close()
	if d.error != nil {
		log.Fatal(err)
		return nil, d.error
	}
	result := []ModelInterface{}
	for d.rows.Next() {
		model := d.modelHandle()
		ptrs := model.Ptrs()
		cols, _ := d.rows.Columns() //返回所有列
		ps := []interface{}{}
		for _, col := range cols {
			ps = append(ps, ptrs[col])
		}
		_ = d.rows.Scan(ps...) //填充数据
		result = append(result, model)
	}
	_ = d.rows.Close()
	if err = d.rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (d *Dao) ToMap() (map[string]interface{}, error) {
	maps, err := d.ToMaps()
	if len(maps) == 0 {
		return nil, err
	}
	return maps[0], err
}

func (d *Dao) ToMaps() ([]map[string]interface{}, error) {
	defer d.rows.Close()
	if d.error != nil {
		return nil, d.error
	}
	//返回所有列
	cols, _ := d.rows.Columns()
	colsLen := len(cols)
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, colsLen)
	//这里表示一行填充数据
	scans := make([]interface{}, colsLen)
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	result := []map[string]interface{}{}
	for d.rows.Next() {
		//填充数据
		d.rows.Scan(scans...)
		row := map[string]interface{}{}
		for k, v := range vals {
			row[cols[k]] = v
		}
		result = append(result, row)
	}
	d.rows.Close()
	if err := d.rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return result, nil
}

func (d *Dao) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	if d.tx != nil {
		return d.tx.Exec(query, args...)
	}
	return d.db.Exec(query, args...)
}

func (d *Dao) Begin() *Dao {
	d.tx, d.error = d.db.Begin()
	return d
}

func (d *Dao) Commit() error {
	if d.tx != nil {
		err := d.tx.Commit()
		d.tx = nil
		return err
	}
	return nil
}

func (d *Dao) Rollback() error {
	if d.tx != nil {
		err := d.tx.Rollback()
		d.tx = nil
		return err
	}
	return nil
}

func (d *Dao) nullType(data interface{}) interface{} {
	switch data.(type) {
	case dtype.NullInt8:
		return data.(dtype.NullInt8).Int8
	case dtype.NullUint8:
		return data.(dtype.NullUint8).Uint8

	case sql.NullInt16:
		return data.(sql.NullInt16).Int16
	case dtype.NullUint16:
		return data.(dtype.NullUint16).Uint16

	case sql.NullInt32:
		return int(data.(sql.NullInt32).Int32)
	case dtype.NullUint32:
		return uint(data.(dtype.NullUint32).Uint32)

	case sql.NullInt64:
		return data.(sql.NullInt64).Int64
	case dtype.NullUint64:
		return data.(dtype.NullUint64).Uint64

	case sql.NullFloat64:
		return data.(sql.NullFloat64).Float64

	case sql.NullString:
		return data.(sql.NullString).String

	case dtype.NullDate:
		return data.(dtype.NullDate).Date
	case dtype.NullTime:
		return data.(dtype.NullTime).Time
	case dtype.NullSet:
		return data.(dtype.NullSet).Set
	case dtype.NullBit:
		return data.(dtype.NullBit).Bit
	case dtype.NullBits:
		return data.(dtype.NullBits).Bits
	}
	return data
}

func (d *Dao) values() (columns []string, values []interface{}, err error) {
	msLend := len(d.data)
	if msLend == 0 {
		return nil, nil, nil
	}
	tags := d.model.tags()
	for column, _ := range tags {
		columns = append(columns, column)
	}
	v := validation.NewValidation()
	for _, model := range d.data {
		model.InsertEvent()
		for _, column := range columns {
			var value interface{}
			if (model.Ass() && model.AssColumn(column)) || tags[column].Default == nil {
				value = model.Get(column)
			} else {
				value = tags[column].Default
			}
			values = append(values, value)
			//检验value数据
			if err = d.checkValues(v, tags[column], column, d.nullType(value)); err != nil {
				return nil, nil, err
			}
		}
	}
	return
}

//checkValues 检查数据是否满足条件
func (d *Dao) checkValues(validation *validation.Validation, tags Tags, column string, value interface{}) error {
	if len(tags.Validators) == 0 {
		return nil
	}
	err := validation.Validator(column, value, tags.Validators, tags.Comment).Error()
	if err != nil {
		return err
	}
	return nil
}


func (d *Dao) sets() (columns []string, values []interface{}, err error) {
	msLend := len(d.data)
	if msLend == 0 {
		return nil, nil, nil
	}
	v := validation.NewValidation()
	tags := d.model.tags()
	model := d.data[0]
	model.UpdateEvent()
	for column, _ := range d.model.AssColumns() {
		columns = append(columns, column)
		value :=  model.Get(column)
		values = append(values, value)
		//检验value数据
		if err = d.checkValues(v, tags[column], column, value); err != nil {
			return nil, nil, err
		}
	}
	return
}

