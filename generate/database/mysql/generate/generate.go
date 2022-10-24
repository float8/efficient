package generate

import (
	"database/sql"
	"github.com/float8/efficient/database"
	"github.com/float8/efficient/generate/database/mysql/dao"
	"github.com/float8/efficient/generate/public"
	"strings"
)

func NewGenerate() *Generate {
	return &Generate{appDirs: map[string]string{}}
}

type Generate struct {
	tables      []string
	columns     []database.ModelInterface
	appDirs     map[string]string
	db          *sql.DB
	dbname      string
	mpkg        string
	dpkg        string
	basePath    string
	projectPath string
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
	modelPath := g.projectPath + g.appDirs["model"]
	code, g.mpkg = newGenerateBaseModel(g.tables, g.dbname, modelPath).generate()
	path := g.absPath("model", "model.go")
	public.WriteFile(path, code)
	return g
}

func (g *Generate) model() *Generate {
	g.baseModel()
	for _, table := range g.tables {
		columns := g.queryColumns(g.dbname, table)
		code := newGenerateModel(table, columns, g.mpkg).generate()
		path := g.absPath("model", strings.ToLower(table)+"_model.go")
		public.WriteFile(path, code)
	}
	return g
}

func (g *Generate) baseDao() *Generate {
	var code string
	daoPath := g.projectPath + g.appDirs["dao"]
	code, g.dpkg = newGenerateBaseDao(g.tables, daoPath).generate()
	path := g.absPath("dao", "dao.go")
	public.WriteFile(path, code)
	return g
}

func (g *Generate) dao() *Generate {
	g.baseDao()
	modelPath := g.projectPath + "/" + g.appDirs["model"]
	for _, table := range g.tables {
		code := newGenerateDao(modelPath, g.dpkg, table).generate()
		path := g.absPath("dao", strings.ToLower(table)+"_dao.go")
		public.WriteFile(path, code)
	}
	return g
}

func (g *Generate) absPath(dirname string, filename string) string {
	return g.basePath + g.projectPath + "/" + g.appDirs[dirname] + "/" + filename
}

func (g *Generate) Execute(db *sql.DB, dbname, basePath, projectPath string, appDirs map[string]string) {
	g.db = db
	g.dbname = dbname
	g.basePath = basePath
	g.projectPath = projectPath
	g.appDirs = appDirs

	g.queryTables().model().dao()
}
