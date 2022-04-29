package model

import (
	"github.com/whf-sky/efficient/widget/database"
)

func NewTables() *Tables {
	tables := &Tables{}
	tables.Init("mysql", "information_schema","_tables" )
	return tables
}

type Tables struct {
	database.Model
	TName string `column:"TABLE_NAME"`
}

func (u *Tables) TableName() string {
	return "tables"
}

func  (u *Tables) Ptrs() map[string]interface{} {
	return map[string]interface{}{
		"TABLE_NAME" : &u.TName,
	}
}

func (c *Tables) Get(key string) interface{} {
	switch key {
	case "TABLE_NAME":
		return c.GetTName()
	}
	return nil
}

func (c *Tables) Set(key string, val interface{}) {
	switch key {
	case "TABLE_NAME":
		c.SetTName(val.(string))
	}
}

func (c *Tables) GetTName() string {
	return c.TName
}

//---------------------set------------------------------

func (c *Tables) SetTName(v string) *Tables {
	c.TName = v
	return c
}