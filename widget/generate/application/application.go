package application

import (
	"github.com/whf-sky/efficient/widget/generate/public"
	"os"
	"path/filepath"
	"strings"
)

func NewApplication(appdirs map[string]string) *Application {
	path, err := filepath.Abs("./")
	if err != nil {
		panic(err)
	}
	return &Application{
		basepath: path,
		appdirs:  appdirs,
	}
}

type Application struct {
	basepath string
	appdirs  map[string]string
}

func (a *Application) Execute() {
	a.dirs()
	a.config()
	a.routers()
	a.service()
	a.main()
}

func (a *Application) dirs() {
	for _, dir := range a.appdirs {
		path := a.basepath + "/" + dir
		err := os.MkdirAll(path, 0766)
		if err != nil {
			panic(err)
		}
	}
}

func (a *Application) config() {
	code := "package config\n\n" +
		"import (\n\t\"github.com/whf-sky/efficient\"\n)\n\n" +
		"func init()  {\n\t" +
		"env := efficient.GetEnv()\n\t" +
		"if env == \"production\" {" +
		"\n\t\tefficient.Config.Addr = \":80\"\n\t\t" +
		"efficient.Config.Debug = true\n\t}\n}\n"
	public.WriteFile(a.basepath+"/"+a.appdirs["config"]+"/config.go", code)
}

func (a *Application) routers() {
	code := "package config\n\n" +
		"import (\n\t\"github.com/whf-sky/efficient\"\n\t\"" +
		"#app_path#/#service_path#\"\n\t\"" +
		"net/http\"\n)\n\n" +
		"func init(){\n\t" +
		"efficient.Routers.Add(\"/test\", &service.TestController{}, http.MethodGet, http.MethodPost)\n" +
		"}"
	apppaths := strings.Split(a.basepath, "/go/src/")
	code = strings.ReplaceAll(code, "#app_path#", apppaths[1])
	code = strings.ReplaceAll(code, "#service_path#", a.appdirs["service"])
	public.WriteFile(a.basepath+"/"+a.appdirs["config"]+"/routers.go", code)
}

func (a *Application) service() {
	code := "package service\n\n" +
		"import \"github.com/whf-sky/efficient\"\n\n" +
		"type TestController struct {\n\tefficient.Controller\n}\n\n" +
		"func (this *TestController) Get(cxt efficient.Context) {\n\tid := cxt.Query(\"id\")\n\tcxt.String(200, \"get:\"+id)\n}\n\n" +
		"func (this *TestController) Post(cxt efficient.Context) {\n\tid := cxt.PostForm(\"id\")\n\tcxt.String(200, \"post:\"+id)\n}"
	public.WriteFile(a.basepath+"/"+a.appdirs["service"]+"/test.go", code)
}

func (a *Application) main() {
	code := "package main\n\n" +
		"import (\n\t\"github.com/whf-sky/efficient\"\n\t" +
		"_ \"#app_path#/config\"\n)\n\n" +
		"func main(){\n\tefficient.WebRun()\n}"
	apppaths := strings.Split(a.basepath, "/go/src/")
	code = strings.ReplaceAll(code, "#app_path#", apppaths[1])
	public.WriteFile(a.basepath+"/"+a.appdirs["cmd"]+"/main.go", code)
}
