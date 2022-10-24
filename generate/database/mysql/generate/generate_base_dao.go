package generate

import (
	"strings"
)

func newGenerateBaseDao(tables []string, path string) *generateBaseDao {
	return &generateBaseDao{
		path:   path,
		tables: tables,
		code: "package #package#" +
			"\n\nimport (\n\t" +
			"\"database/sql\"\n\t" +
			"_ \"github.com/go-sql-driver/mysql\"\n\t" +
			"\"github.com/float8/efficient.demo/config\"\n\t" +
			"\"github.com/float8/efficient/database\"\n)\n\n" +
			"var db = func() *sql.DB {\n\t" +
			"return database.NewDb().MysqlDsn(\n\t\t" +
			"config.DbConfig.Addr,\n\t\t" +
			"config.DbConfig.Port,\n\t\t" +
			"config.DbConfig.Account,\n\t\t" +
			"config.DbConfig.Passwd,\n\t\t" +
			"config.DbConfig.Dbname,\n\t\t" +
			"config.DbConfig.Charset,\n\t" +
			").Open(func(db *sql.DB) {\n\t\t" +
			"db.SetConnMaxLifetime(config.DbConfig.SetConnMaxLifetime) //设置最大连接生存时间\n\t\t" +
			"db.SetConnMaxIdleTime(config.DbConfig.ConnMaxIdleTime) //设置连接最大空闲时间\n\t\t" +
			"db.SetMaxOpenConns(config.DbConfig.MaxOpenConns) //设置最大打开链接数\n\t\t" +
			"db.SetMaxIdleConns(config.DbConfig.MaxIdleConns) //设置最大空闲链接\n\t})\n}()\n\n" +
			"type Dao struct {\n\tdatabase.Dao\n}\n\n" +
			"func (d *Dao) Init(fun func() database.ModelInterface) *Dao {\n\td.SetDb(\"mysql\", db)\n\td.SetModel(fun)\n\treturn d\n}\n",
	}
}

type generateBaseDao struct {
	code   string
	path   string
	tables []string
	pkg    string
}

func (g *generateBaseDao) generate() (code, pkg string) {
	path := strings.Trim(g.path, "/")
	paths := strings.Split(path, "/")
	g.pkg = paths[len(paths)-1]
	return strings.ReplaceAll(g.code, "#package#", g.pkg), g.pkg
}
