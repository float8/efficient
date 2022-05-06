package generate

import (
	"strings"
)

func newGenerateBaseDao(tables []string, path string) *generateBaseDao {
	return &generateBaseDao{
		path:   path,
		tables: tables,
		code: "package #package#\n\nimport (\n\t\"database/sql\"\n\t_ \"github.com/go-sql-driver/mysql\"\n\t\"github.com/whf-sky/efficient/widget/database\"\n\t\"time\"\n)\n\n" +
			"//请重新配置数据连接信息\nvar db = func() *sql.DB {\n\treturn database.NewDb().MysqlDsn(\"127.0.0.1\", \"3306\", \"root\", \"123456\", \"test\", \"utf8mb4\").Open(func(db *sql.DB) {\n\t\t" +
			"db.SetConnMaxIdleTime(time.Minute * 4) //设置最大连接生存时间\n\t\t" +
			"db.SetMaxOpenConns(5)                  //设置最大打开链接数\n\t\t" +
			"db.SetMaxIdleConns(5)                  //设置最大空闲链接\n\t\t" +
			"db.SetConnMaxLifetime(time.Minute * 2) //设置连接最大空闲时间\n\t})\n}()\n\n" +
			"type Dao struct {\n\tdatabase.Dao\n}\n\nfunc (d *Dao) Init() *Dao {\n\td.SetDb(db)\n\td.DriverName(\"mysql\")\n\treturn d\n}\n",
	}
}

type generateBaseDao struct {
	code string
	path          string
	tables        []string
	pkg string
}

func (g *generateBaseDao) generate() (code, pkg string) {
	path := strings.Trim(g.path, "/")
	paths := strings.Split(path, "/")
	g.pkg = paths[len(paths)-1]
	return strings.ReplaceAll(g.code, "#package#", g.pkg), g.pkg
}


