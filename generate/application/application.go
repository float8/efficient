package application

import (
	"github.com/float8/efficient/generate/public"
	"os"
	"strings"
)

func NewApplication(basePath, projectPath string, appDirs map[string]string) *Application {
	return &Application{
		basePath:    basePath,
		projectPath: projectPath,
		appDirs:     appDirs,
	}
}

type Application struct {
	basePath    string
	projectPath string
	appDirs     map[string]string
}

func (a *Application) absPath(dirname, file string) string {
	return a.basePath + a.projectPath + "/" + a.appDirs[dirname] + "/" + file
}

func (a *Application) Execute() {
	a.dirs()
	a.config()
	a.loggerConfig()
	a.databaseConfig()
	a.routers()
	a.service()
	a.main()
}

func (a *Application) dirs() {
	for _, dir := range a.appDirs {
		path := a.basePath + a.projectPath + "/" + dir
		err := os.MkdirAll(path, 0766)
		if err != nil {
			panic(err)
		}
	}
}

func (a *Application) config() {
	code := "package config\n\n" +
		"import (\n\t\"github.com/float8/efficient\"\n)\n\n" +
		"func init()  {\n\t" +
		"env := efficient.GetEnv()\n\t" +
		"if env == \"production\" {" +
		"\n\t\tefficient.Config.Addr = \":80\"\n\t\t" +
		"efficient.Config.Debug = true\n\t}\n}\n"
	absPath := a.absPath("config", "config.go")
	public.WriteFile(absPath, code)
}

func (a *Application) loggerConfig() {
	code := "package config\n\n" +
		"import (\n\t\"" +
		"github.com/gin-gonic/gin\"\n\t\"" +
		"github.com/sirupsen/logrus\"\n\t\"" +
		"github.com/float8/efficient\"\n\t\"" +
		"os\"\n)\n\n" +
		"func init()  {\n\t" +
		"env := efficient.GetEnv()\n\t" +
		"if env == \"production\" {\n\t\t" +
		"gin.DefaultWriter = os.Stdout\n\t\t" +
		"gin.DefaultErrorWriter = os.Stderr\n\t\t" +
		"efficient.SetLogger(func(logger *logrus.Logger, log *logrus.Entry) {\n\t\t\t" +
		"logger.SetReportCaller(false)\n\t\t\t" +
		"log = logrus.NewEntry(logger)\n\t\t\t" +
		"logger.Out = os.Stdout\n\t\t})\n\t}\n}"

	absPath := a.absPath("config", "logger.go")
	public.WriteFile(absPath, code)
}

func (a *Application) databaseConfig() {
	code := "package config\n\n" +
		"import (\n\t\"github.com/float8/efficient\"\n\t\"time\"\n)\n\n" +
		"var DbConfig dbConfig\n\n" +
		"type dbConfig struct {\n\t" +
		"driver             string\n\t" +
		"Addr               string\n\t" +
		"Port               string\n\t" +
		"Account            string\n\t" +
		"Passwd             string\n\t" +
		"Dbname             string\n\t" +
		"Charset            string\n\t" +
		"ConnMaxIdleTime    time.Duration\n\t" +
		"SetConnMaxLifetime time.Duration\n\t" +
		"MaxOpenConns       int\n\t" +
		"MaxIdleConns       int\n}\n\n" +
		"func init() {\n\t" +
		"env := efficient.GetEnv()\n\t" +
		"if env == \"production\" {\n\t\t" +
		"DbConfig = dbConfig{\n\t\t\tA" +
		"ddr:               \"127.0.0.1\",\n\t\t\t" +
		"Port:               \"3306\",\n\t\t\t" +
		"Account:            \"root\",\n\t\t\t" +
		"Passwd:             \"123456\",\n\t\t\t" +
		"Dbname:             \"test\",\n\t\t\t" +
		"Charset:            \"utf8mb4\",\n\t\t\t" +
		"SetConnMaxLifetime: time.Minute * 4,\n\t\t\t" +
		"ConnMaxIdleTime:    time.Minute * 2,\n\t\t\t" +
		"MaxOpenConns:       5,\n\t\t\t" +
		"MaxIdleConns:       5,\n\t\t}\n\t}\n}\n"
	absPath := a.absPath("config", "database.go")
	public.WriteFile(absPath, code)
}

func (a *Application) routers() {
	code := "package config\n\n" +
		"import (\n\t\"github.com/float8/efficient\"\n\t\"" +
		"#app_path#/#service_path#\"\n\t\"" +
		"net/http\"\n)\n\n" +
		"func init(){\n\t" +
		"efficient.Routers.Add(\"/test\", &service.TestController{}, http.MethodGet, http.MethodPost)\n" +
		"}"

	code = strings.ReplaceAll(code, "#app_path#", a.projectPath)
	code = strings.ReplaceAll(code, "#service_path#", a.appDirs["service"])
	absPath := a.absPath("config", "routers.go")
	public.WriteFile(absPath, code)
}

func (a *Application) service() {
	code := "package service\n\n" +
		"import \"github.com/float8/efficient\"\n\n" +
		"type TestController struct {\n\tefficient.Controller\n}\n\n" +
		"func (this *TestController) Get(cxt efficient.Context) {\n\tid := cxt.Query(\"id\")\n\tcxt.String(200, \"get:\"+id)\n}\n\n" +
		"func (this *TestController) Post(cxt efficient.Context) {\n\tid := cxt.PostForm(\"id\")\n\tcxt.String(200, \"post:\"+id)\n}"

	absPath := a.absPath("service", "test.go")
	public.WriteFile(absPath, code)
}

func (a *Application) main() {
	code := "package main\n\n" +
		"import (\n\t\"github.com/float8/efficient\"\n\t" +
		"_ \"#app_path#/config\"\n)\n\n" +
		"func main(){\n\tefficient.WebRun()\n}"
	code = strings.ReplaceAll(code, "#app_path#", a.projectPath)

	absPath := a.absPath("cmd", "main.go")
	public.WriteFile(absPath, code)
}
