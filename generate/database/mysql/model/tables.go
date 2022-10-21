package model

import (
	"github.com/float8/efficient/database"
)

func NewTables() *Tables {
	tables := &Tables{}
	tables.Init("mysql", "information_schema", "_tables")
	return tables
}

type Tables struct {
	database.Model
	TName string `column:"TABLE_NAME"`
}

func (t *Tables) TableName() string {
	return "tables"
}

func (t *Tables) Ptrs() map[string]interface{} {
	return map[string]interface{}{
		"TABLE_NAME": &t.TName,
	}
}

func (t *Tables) Get(key string) interface{} {
	switch key {
	case "TABLE_NAME":
		return t.GetTName()
	}
	return nil
}

func (t *Tables) Set(key string, val interface{}) {
	switch key {
	case "TABLE_NAME":
		t.SetTName(val.(string))
	}
}

func (t *Tables) GetTName() string {
	return t.TName
}

//---------------------set------------------------------

func (t *Tables) SetTName(v string) *Tables {
	t.TName = v
	return t
}
