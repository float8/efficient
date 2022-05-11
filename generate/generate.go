package generate

import (
	"database/sql"
	"github.com/whf-sky/efficient/generate/application"
	"github.com/whf-sky/efficient/generate/database/mysql/generate"
	"log"
	"path/filepath"
)

var generates = map[string]GenerateInterface{
	"mysql": generate.NewGenerate(),
}

func RegisterGenerate(DriverName string, generate GenerateInterface) {
	generates[DriverName] = generate
}

type GenerateInterface interface {
	Execute(sqlDb *sql.DB, dbname, modelPath, daoPath string)
}

func NewGenerate() *Generate {
	basePath, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	return &Generate{
		basePath: basePath,
		appdirs: map[string]string{
			"dao":     "application/dao",
			"model":   "application/model",
			"service": "application/service",
			"config":  "config",
			"cmd":     "cmd",
		},
		modelPath: basePath + "/application/model",
		daoPath:   basePath + "/application/dao",
	}
}

type Generate struct {
	Db         *sql.DB
	DriverName string
	basePath   string
	modelPath  string
	daoPath    string
	appdirs    map[string]string
}

func (g *Generate) SetAppDir(appdirs map[string]string) *Generate {
	g.appdirs = appdirs
	return g
}

func (g *Generate) Application() *Generate {
	application.NewApplication(g.appdirs).Execute()
	return g
}

func (g *Generate) SetDb(driverName string, db *sql.DB) *Generate {
	g.Db = db
	g.DriverName = driverName
	return g
}

func (g *Generate) Service(path string) *Generate {

	return g
}

func (g *Generate) Database(dbname string) *Generate {
	generates[g.DriverName].Execute(g.Db, dbname, g.modelPath, g.daoPath)
	log.Println("File generation complete！")
	return g
}
