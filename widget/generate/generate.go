package generate

import (
	"database/sql"
	generate2 "github.com/whf-sky/efficient/widget/generate/database/mysql/generate"
	"log"
	"path/filepath"
)

var generates = map[string]GenerateInterface{
	"mysql": generate2.NewGenerate(),
}

func RegisterGenerate(key string, generate GenerateInterface) {
	generates[key] = generate
}

type GenerateInterface interface {
	Execute(sqlDb *sql.DB, dbname, modelPath, daoPath string)
}

func NewGenerate(driverName string, db *sql.DB) *Generate {
	return &Generate{
		Db:         db,
		DriverName: driverName,
	}
}

type Generate struct {
	Db         *sql.DB
	DriverName string
}

func (g *Generate) Service(path string) {

}

func (g *Generate) Database(dbname, modelPath, daoPath string) {
	appPath, _ := filepath.Abs("./")
	generates[g.DriverName].Execute(g.Db, dbname, appPath+"/"+modelPath, appPath+"/"+daoPath)
	log.Println("File generation complete！")
}
