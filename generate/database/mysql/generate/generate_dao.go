package generate

import (
	estrings "github.com/whf-sky/efficient/tools/strings"
	"strings"
)

func newGenerateDao(modelPath , pkg, tableName string) *generateDao {
	return &generateDao{
		pkg:       pkg,
		tableName: tableName,
		modelPath: modelPath,
		sections: map[string]string{
			"package":   "package #package#",
			"import":    "import (\n\t\"#modelimport#\"\n\t\"github.com/whf-sky/efficient/database\"\n)",
			"new":       "func New#model#Dao() *#model#Dao {\n\td := &#model#Dao{}\n\td.Init(func() database.ModelInterface {return model.New#model#()})\n\treturn d\n}",
			"struct":   "type #model#Dao struct {\n\tDao\n}",
		},
		sectionsOrder:     []string{"package", "import", "new", "struct"},
	}
}
type generateDao struct {
	tableName string
	modelPath string
	pkg string
	sections map[string]string
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
	modelPath := strings.Split(g.modelPath, "go/src/")
	g.sections["import"] = strings.ReplaceAll(g.sections["import"], "#modelimport#", modelPath[1])
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
