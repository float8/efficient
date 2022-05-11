package database

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/whf-sky/efficient/validation"
)

func nullType(data interface{}) interface{} {
	switch data.(type) {
	case NullInt8:
		return data.(NullInt8).Int8
	case NullUint8:
		return data.(NullUint8).Uint8

	case sql.NullInt16:
		return data.(sql.NullInt16).Int16
	case NullUint16:
		return data.(NullUint16).Uint16

	case sql.NullInt32:
		return int(data.(sql.NullInt32).Int32)
	case NullUint32:
		return uint(data.(NullUint32).Uint32)

	case sql.NullInt64:
		return data.(sql.NullInt64).Int64
	case NullUint64:
		return data.(NullUint64).Uint64

	case sql.NullFloat64:
		return data.(sql.NullFloat64).Float64

	case sql.NullString:
		return data.(sql.NullString).String

	case NullDate:
		return data.(NullDate).Date
	case NullTime:
		return data.(NullTime).Time
	case NullSet:
		return data.(NullSet).Set
	case NullBit:
		return data.(NullBit).Bit
	case NullBits:
		return data.(NullBits).Bits
	}
	return data
}

func values(d *Dao) (kvs []map[string]interface{}, columns []string, values []interface{}, err error) {
	if len(d.data) == 0 {
		err = errors.New("insert data does not exist")
		return
	}
	tags := d.model.tags()
	for column := range tags {
		columns = append(columns, column)
	}
	v := validation.NewValidation()
	for _, model := range d.data {
		model.InsertEvent()
		kv := map[string]interface{}{}
		for _, column := range columns {
			var value interface{}
			if (model.Ass() && model.AssColumn(column)) || tags[column].Default == nil {
				value = model.Get(column)
			} else {
				value = tags[column].Default
			}
			kv[column] = value
			values = append(values, value)
			if err == nil {
				err = checkValues(v, tags[column], column, nullType(value))
			}
		}
		kvs = append(kvs, kv)
	}
	return
}

func sets(d *Dao) (kv map[string]interface{}, columns []string, values []interface{}, err error) {
	if len(d.data) == 0 {
		err = errors.New("update data does not exist")
		return
	}
	v := validation.NewValidation()
	tags := d.model.tags()
	model := d.data[0]
	model.UpdateEvent()
	kv = map[string]interface{}{}
	for column := range model.AssColumns() {
		columns = append(columns, column)
		value := model.Get(column)
		values = append(values, value)
		kv[column] = value
		if err == nil {
			err = checkValues(v, tags[column], column, value)
		}
	}
	return
}

//checkValues 检查数据是否满足条件
func checkValues(validation *validation.Validation, tags Tags, column string, value interface{}) error {
	if len(tags.Validators) == 0 {
		return nil
	}
	err := validation.Validator(column, value, tags.Validators, tags.Comment).Error()
	if err != nil {
		return err
	}
	return nil
}

func jsonMarshal(data interface{}) string {
	str, _ := json.Marshal(data)
	return string(str)
}
