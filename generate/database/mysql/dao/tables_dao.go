package dao

import (
	"github.com/whf-sky/efficient/database"
	"github.com/whf-sky/efficient/generate/database/mysql/model"
)

func NewTablesDao() *TablesDao {
	userDao := &TablesDao{}
	userDao.SetModel(func() database.ModelInterface { return model.NewTables() })
	return userDao
}

type TablesDao struct {
	database.Dao
}

func (t *TablesDao) QueryTables(dbname string) []string {
	tables, err := t.Query("select TABLE_NAME from information_schema.tables where table_schema = '" + dbname + "'").ToModels()
	if err != nil {
		panic(err)
	}

	tbls := []string{}
	for _, table := range tables {
		m := table.(*model.Tables)
		tbls = append(tbls, m.GetTName())
	}
	return tbls
}
