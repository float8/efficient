package database

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

func newQuery(dao *Dao, log *logrus.Entry) *query  {
	return &query{dao: dao, log: log}
}

type query struct {
	dao *Dao
	log *logrus.Entry
}

func (q *query) Rows() (*sql.Rows, error) {
	if q.dao.error != nil {
		q.log.Error(q.dao.error)
	}
	return q.dao.rows, q.dao.error
}

func (q *query) ToModel() (model ModelInterface, err error) {
	models, err := q.ToModels()
	if len(models) == 0 {
		return nil, err
	}
	return models[0], err
}

func (q *query) ToModels() (models []ModelInterface, err error) {
	defer func() {
		err := q.dao.rows.Close()
		if err != nil {
			q.log.Error(err)
		}
	}()
	if q.dao.error != nil {
		q.log.Error(q.dao.error)
		return nil, q.dao.error
	}
	result := []ModelInterface{}
	for q.dao.rows.Next() {
		model := q.dao.modelHandle()
		ptrs := model.Ptrs()
		cols, err := q.dao.rows.Columns() //返回所有列
		if err != nil {
			q.log.Error(err)
			return nil, err
		}
		ps := []interface{}{}
		for _, col := range cols {
			ps = append(ps, ptrs[col])
		}
		err = q.dao.rows.Scan(ps...) //填充数据
		if err != nil {
			q.log.Error(err)
			return nil, err
		}
		result = append(result, model)
	}
	err = q.dao.rows.Close()
	if err != nil {
		q.log.Error(err)
	}
	if err = q.dao.rows.Err(); err != nil {
		q.log.Error(err)
		return nil, err
	}
	q.log.Info(jsonMarshal(result))
	return result, nil
}

func (q *query) ToMap() (result map[string]interface{}, err error) {
	maps, err := q.ToMaps()
	if len(maps) == 0 {
		return nil, err
	}
	return maps[0], err
}

func (q *query) ToMaps() (result []map[string]interface{}, err error) {
	defer func() {
		err := q.dao.rows.Close()
		if err != nil {
			q.log.Error(err)
		}
	}()
	if q.dao.error != nil {
		q.log.Error(q.dao.error)
		return nil, q.dao.error
	}
	//返回所有列
	cols, _ := q.dao.rows.Columns()
	colsLen := len(cols)
	//这里表示一行所有列的值，用[]byte表示
	vals := make([][]byte, colsLen)
	//这里表示一行填充数据
	scans := make([]interface{}, colsLen)
	//这里scans引用vals，把数据填充到[]byte里
	for k, _ := range vals {
		scans[k] = &vals[k]
	}
	result = []map[string]interface{}{}
	for q.dao.rows.Next() {
		//填充数据
		err := q.dao.rows.Scan(scans...)
		if err != nil {
			q.log.Error(err)
			return nil, err
		}
		row := map[string]interface{}{}
		for k, v := range vals {
			row[cols[k]] = v
		}
		result = append(result, row)
	}
	err = q.dao.rows.Close()
	if err != nil {
		q.log.Error(err)
		return
	}

	if err = q.dao.rows.Err(); err != nil {
		q.log.Error(err)
		return
	}
	q.log.Info(jsonMarshal(result))
	return result, nil
}