package dao

import (
	"github.com/whf-sky/efficient/database"
	"github.com/whf-sky/efficient/generate/database/mysql/model"
)

func NewColumnsDao() *ColumnsDao {
	dao := &ColumnsDao{}
	dao.SetModel(func() database.ModelInterface { return model.NewColumns() })
	return dao
}

type ColumnsDao struct {
	database.Dao
}

func (c *ColumnsDao) QueryColumns(database, table_name string) []database.ModelInterface {
	var query = "select * from information_schema.columns " +
		"where table_schema= '" + database + "' and table_name='" + table_name + "' " +
		"order by table_name, ordinal_position asc"

	models, err := c.Query(query).ToModels()
	if err != nil {
		panic(err)
	}
	return models
}
