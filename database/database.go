package database

import (
	"database/sql"
	"fmt"
	"log"
)

func NewDb() *Db {
	return &Db{}
}

type Db struct {
	db         *sql.DB
	dsn        string
	driverName string
}

func (d *Db) MysqlDsn(addr, port, account, passwd, dbname, charset string) *Db {
	d.driverName = "mysql"
	d.dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
		account,
		passwd,
		addr,
		port,
		dbname,
		charset)
	return d
}

//Open 打开数据库连接
//d.SetConnMaxLifetime(time.Minute * 4) 设置最大连接生存时间
//d.SetMaxOpenConns(5) 设置最大打开链接数
//d.SetMaxIdleConns(5) 设置最大空闲链接
//d.SetConnMaxIdleTime(time.Minute * 2) 设置连接最大空闲时间
func (d *Db) Open(sets ...func(db *sql.DB)) *sql.DB {
	var err error
	d.db, err = sql.Open(d.driverName, d.dsn)
	if err != nil {
		log.Println("dataSourceName: " + d.dsn)
		panic("数据源配置不正确: " + err.Error())
	}

	if len(sets) > 0 {
		sets[0](d.db)
	}

	if err = d.db.Ping(); nil != err {
		panic("数据库链接失败: " + err.Error())
	}
	return d.db
}
