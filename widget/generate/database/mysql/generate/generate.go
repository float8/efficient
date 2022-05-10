package generate

import (
	"database/sql"
	"github.com/whf-sky/efficient/widget/database"
	"github.com/whf-sky/efficient/widget/generate/database/mysql/dao"
	"github.com/whf-sky/efficient/widget/generate/public"
	"strings"
)

func NewGenerate() *Generate {
	return &Generate{}
}

type Generate struct {
	tables    []string
	columns   []database.ModelInterface
	modelPath string
	daoPath   string
	db        *sql.DB
	dbname    string
	mpkg      string
	dpkg      string
}

func (g *Generate) queryTables() *Generate {
	tablesDao := dao.NewTablesDao()
	tablesDao.SetDb("mysql", g.db)
	g.tables = tablesDao.QueryTables(g.dbname)
	return g
}

func (g *Generate) queryColumns(dbname string, tablename string) []database.ModelInterface {
	columnsDao := dao.NewColumnsDao()
	columnsDao.SetDb("mysql", g.db)
	return columnsDao.QueryColumns(dbname, tablename)
}

func (g *Generate) baseModel() *Generate {
	var code string
	code, g.mpkg = newGenerateBaseModel(g.tables, g.dbname, g.modelPath).generate()
	path := g.modelPath + "/model.go"
	public.WriteFile(path, code)
	return g
}

func (g *Generate) model() *Generate {
	g.baseModel()
	for _, table := range g.tables {
		columns := g.queryColumns(g.dbname, table)
		code := newGenerateModel(table, columns, g.mpkg).generate()
		path := g.modelPath + "/" + strings.ToLower(table) + "_model.go"
		public.WriteFile(path, code)
	}
	return g
}

func (g *Generate) baseDao() *Generate {
	var code string
	code, g.dpkg = newGenerateBaseDao(g.tables, g.daoPath).generate()
	path := g.daoPath + "/dao.go"
	public.WriteFile(path, code)
	return g
}

func (g *Generate) dao() *Generate {
	g.baseDao()
	for _, table := range g.tables {
		code := newGenerateDao(g.modelPath, g.dpkg, table).generate()
		path := g.daoPath + "/" + strings.ToLower(table) + "_dao.go"
		public.WriteFile(path, code)
	}
	return g
}

func (g *Generate) Execute(db *sql.DB, dbname, modelPath, daoPath string) {
	g.daoPath = daoPath
	g.modelPath = modelPath
	g.dbname = dbname
	g.db = db
	g.queryTables().model().dao()
}


