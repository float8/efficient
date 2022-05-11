package database

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
)

type ModelHandle func() ModelInterface

type DaoInterface interface{}

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
	error       error
}

func (d *Dao) SetDb(driverName string, db *sql.DB) *sql.DB {
	d.db = db
	d.driverName = driverName
	return d.db
}

func (d *Dao) SetModel(fn func() ModelInterface) *Dao {
	d.modelHandle = fn
	d.model = fn()
	return d
}

func (d *Dao) Model() ModelInterface {
	return d.model
}

func (d *Dao) TableName(indexs ...int) string {
	if len(d.data) > 0 {
		index := 0
		if len(indexs) > 0 {
			index = indexs[0]
		}
		return d.data[index].TableName()
	}
	return d.model.TableName()
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

func (d *Dao) Insert() (result sql.Result, err error) {
	//获取数据并做规则验证
	kvs, columns, values, err := values(d)
	log := Log.WithFields(logrus.Fields{
		"key":  "[SQL][Insert]",
		"data": jsonMarshal(kvs),
	})
	if err != nil {
		log.Debug(err)
		return nil, err
	}

	//生成Insert SQL
	vLen := len(values)
	sql, err := SQL(d.driverName).InsertQuery(d.TableName(), columns, vLen)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log = log.WithField("sql", sql)

	//执行SQL
	result, err = d.Exec(sql, values...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	log.Info("ok")
	return
}

func (d *Dao) Update() (rowsAffected int64, err error) {
	//处理修改数据
	kv, columns, values, err := sets(d)
	log := Log.WithFields(logrus.Fields{
		"key":  "[SQL][Update]",
		"data": jsonMarshal(kv),
	})
	if err != nil {
		log.Debug(err)
		return
	}

	//生成SQL
	sql, err := SQL(d.driverName).UpdateQuery(d.TableName(), columns, d.where)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	log = log.WithField("sql", sql)

	//执行SQL
	args := append(values, d.whereArgs...)
	result, err := d.Exec(sql, args...)
	if err != nil {
		log.Error(err)
		return
	}

	//影响行数
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error(err)
		return
	}
	log.WithField("rows_affected", rowsAffected).Info("ok")
	return
}

func (d *Dao) Delete() (rowsAffected int64, err error) {
	sql, _ := SQL(d.driverName).DeleteQuery(d.TableName(), d.where)
	result, err := d.Exec(sql, d.whereArgs...)
	log := Log.WithFields(logrus.Fields{
		"key":  "[SQL][Delete]",
		"data": jsonMarshal(d.whereArgs),
		"sql":  sql,
	})
	if err != nil {
		log.Error(err)
		return
	}

	//影响函数
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		log.Error(err)
		return
	}
	log.WithField("rows_affected", rowsAffected).Info("ok")
	return
}

func (d *Dao) QueryRow(query string, args ...interface{}) *sql.Row {
	query, _ = SQL(d.driverName).SelectQuery(d.TableName(), query)
	Log.WithFields(logrus.Fields{
		"key":  "[SQL][QueryRow]",
		"data": jsonMarshal(args),
		"sql":  query,
	}).Info("ok")
	if d.tx != nil {
		return d.tx.QueryRow(query, args...)
	}
	return d.db.QueryRow(query, args...)
}

func (d *Dao) Query(query string, args ...interface{}) *query {
	query, _ = SQL(d.driverName).SelectQuery(d.TableName(), query)
	fmt.Println(Log)
	log := Log.WithFields(logrus.Fields{
		"key":  "[SQL][Query]",
		"data": jsonMarshal(args),
		"sql":  query,
	})
	if d.tx == nil {
		d.rows, d.error = d.db.Query(query, args...)
	} else {
		d.rows, d.error = d.tx.Query(query, args...)
	}
	return newQuery(d, log)
}

func (d *Dao) Exec(query string, args ...interface{}) (result sql.Result, err error) {
	if d.tx != nil {
		return d.tx.Exec(query, args...)
	}
	return d.db.Exec(query, args...)
}

func (d *Dao) Begin() error {
	d.tx, d.error = d.db.Begin()
	return d.error
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
