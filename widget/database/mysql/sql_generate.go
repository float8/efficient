package mysql

import (
	"errors"
	"strings"
)

type Mysql struct {
	insertQuery string
	deleteQuery string
	updateQuery string
}

func NewMysql() *Mysql {
	return &Mysql{
		insertQuery: "insert into `#table#` #columns# values #vals#",
		deleteQuery: "delete from `#table#`#where#",
		updateQuery: "update `#table#` set #sets##where#",
	}
}

func (m *Mysql) InsertQuery(table string, columns []string, valsLen int) (sql string, err error) {
	clmsLen := len(columns)
	if clmsLen == 0 {
		return "", errors.New("The columns values does not exist!")
	}
	vals := "(" + strings.Repeat("?,", clmsLen-1) + "?)"
	m.insertQuery = strings.Replace(m.insertQuery, "#table#", table, 1)
	m.insertQuery = strings.Replace(m.insertQuery, "#columns#", "(`"+strings.Join(columns, "`,`")+"`)", 1)
	m.insertQuery = strings.Replace(m.insertQuery, "#vals#", strings.Repeat(vals+",", valsLen/clmsLen-1)+vals, 1)
	return m.insertQuery, nil
}

func (m *Mysql) UpdateQuery(table string, sets []string, where string) (sql string, err error) {
	slen := len(sets)
	if slen == 0 {
		return "", errors.New("The set does not exist!")
	}
	if where != "" {
		where = " where " + where
	}
	m.updateQuery = strings.Replace(m.updateQuery, "#table#", table, 1)
	m.updateQuery = strings.Replace(m.updateQuery, "#where#", where, 1)
	m.updateQuery = strings.Replace(m.updateQuery, "#sets#", strings.Join(sets, "=?,")+"=?", 1)
	return m.updateQuery, nil
}

func (m *Mysql) DeleteQuery(table string, where string) (sql string, err error) {
	if where != "" {
		where = " where " + where
	}
	m.deleteQuery = strings.Replace(m.deleteQuery, "#table#", table, 1)
	m.deleteQuery = strings.Replace(m.deleteQuery, "#where#", where, 1)
	return m.deleteQuery, nil
}
