package generate

import (
	"database/sql"
	"github.com/float8/efficient/generate/application"
	"github.com/float8/efficient/generate/database/mysql/generate"
	"github.com/float8/efficient/generate/public"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var generates = map[string]GenerateInterface{
	"mysql": generate.NewGenerate(),
}

func RegisterGenerate(DriverName string, generate GenerateInterface) {
	generates[DriverName] = generate
}

type GenerateInterface interface {
	Execute(sqlDb *sql.DB, dbname, basePath, projectPath string, appDirs map[string]string)
}

func NewGenerate() *Generate {
	basePath, err := filepath.Abs("./")
	basePath = strings.ReplaceAll(basePath, "\\", "/")
	if err != nil {
		panic(err)
	}
	projectPath := projectPath(basePath)

	basePath = strings.Replace(basePath, projectPath, "", 1)

	appDirs := map[string]string{
		"dao":     "application/dao",
		"model":   "application/model",
		"service": "application/service",
		"config":  "config",
		"cmd":     "cmd",
	}
	return &Generate{
		appDirs:     appDirs,
		basePath:    basePath,
		projectPath: projectPath,
	}
}

func projectPath(basePath string) string {
	GO111MODULE := os.Getenv("GO111MODULE")
	ok, _ := public.PathExists(basePath + "/go.mod")
	if GO111MODULE == "off" || (GO111MODULE == "auto" && !ok) {
		GOPATH := os.Getenv("GOPATH")
		return basePath[len(GOPATH)+5:]
	}
	if !ok {
		panic("The go.mod file is missing！")
	}
	module := public.ReadFirstLine(basePath + "/go.mod")
	return strings.TrimLeft(module, "module ")
}

type Generate struct {
	Db          *sql.DB
	DriverName  string
	basePath    string
	projectPath string
	appDirs     map[string]string
}

func (g *Generate) SetAppDir(appDirs map[string]string) *Generate {
	g.appDirs = appDirs
	return g
}

func (g *Generate) Application() *Generate {
	application.NewApplication(g.basePath, g.projectPath, g.appDirs).Execute()
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
	generates[g.DriverName].Execute(g.Db, dbname, g.basePath, g.projectPath, g.appDirs)
	log.Println("File generation complete！")
	return g
}
