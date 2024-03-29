package generate

import (
	estrings "github.com/float8/efficient/tools/strings"
	"strings"
)

func newGenerateDao(modelPath, pkg, tableName string) *generateDao {
	return &generateDao{
		pkg:       pkg,
		tableName: tableName,
		modelPath: modelPath,
		sections: map[string]string{
			"package": "package #package#",
			"import":  "import (\n\t\"#modelimport#\"\n\t\"github.com/float8/efficient/database\"\n)",
			"new":     "func New#model#Dao() *#model#Dao {\n\td := &#model#Dao{}\n\td.Init(func() database.ModelInterface {return model.New#model#()})\n\treturn d\n}",
			"struct":  "type #model#Dao struct {\n\tDao\n}",
		},
		sectionsOrder: []string{"package", "import", "new", "struct"},
	}
}

type generateDao struct {
	tableName     string
	modelPath     string
	pkg           string
	sections      map[string]string
	sectionsOrder []string
}

func (g *generateDao) generate() string {
	g.packageHandle()
	g.importHandle()
	g.modelHandle()
	code := []string{}
	for _, key := range g.sectionsOrder {
		code = append(code, g.sections[key])
	}
	return strings.Join(code, "\n\n")
}

func (g *generateDao) packageHandle() {
	g.sections["package"] = strings.ReplaceAll(g.sections["package"], "#package#", g.pkg)
}

func (g *generateDao) importHandle() {
	g.sections["import"] = strings.ReplaceAll(g.sections["import"], "#modelimport#", g.modelPath)
}

func (g *generateDao) modelHandle() {
	modelSlice := strings.Split(g.tableName, "_")
	for i, s := range modelSlice {
		modelSlice[i] = estrings.FirstUpper(s)
	}
	model := strings.Join(modelSlice, "")
	for key, section := range g.sections {
		g.sections[key] = strings.ReplaceAll(section, "#model#", model)
	}
}
