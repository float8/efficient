package generate

import (
	estrings "github.com/float8/efficient/tools/strings"
	"strings"
)

func newGenerateBaseModel(tables []string, dbname, path string) *generateBaseModel {
	return &generateBaseModel{
		path:   path,
		dbname: dbname,
		tables: tables,
		sections: map[string]string{
			"package": "package #package#",
			"import":  "import \"github.com/float8/efficient/database\"",
			"struct":  "type Model struct {\n\tdatabase.Model\n}",
			"minit":   "func (m *Model) init(key string) *Model {\n\tm.Model.Init(\"mysql\",\"#dbname#\", key)\n\treturn m\n}",
			"init":    "func init() {\n#register#\n}",
		},
		sectionsOrder: []string{"package", "import", "struct", "minit", "init"},
	}
}

type generateBaseModel struct {
	dbname        string
	path          string
	tables        []string
	sections      map[string]string
	sectionsOrder []string
	pkg           string
}

func (g *generateBaseModel) generate() (code, pkg string) {
	g.packageHandle()
	g.mInitHandle()
	g.initHandle()
	codes := []string{}
	for _, key := range g.sectionsOrder {
		codes = append(codes, g.sections[key])
	}
	return strings.Join(codes, "\n\n"), g.pkg
}

func (g *generateBaseModel) packageHandle() {
	path := strings.Trim(g.path, "/")
	paths := strings.Split(path, "/")
	g.pkg = paths[len(paths)-1]
	g.sections["package"] = strings.ReplaceAll(g.sections["package"], "#package#", g.pkg)
}

func (g *generateBaseModel) mInitHandle() {
	g.sections["minit"] = strings.ReplaceAll(g.sections["minit"], "#dbname#", g.dbname)
}

func (g *generateBaseModel) initHandle() {
	register := []string{}
	for _, table := range g.tables {
		tpl := "\tdatabase.RegisterModel(New#model#())"
		tpl = strings.ReplaceAll(tpl, "#model#", g.model(table))
		register = append(register, tpl)
	}
	g.sections["init"] = strings.ReplaceAll(g.sections["init"], "#register#", strings.Join(register, "\n"))
}

func (g *generateBaseModel) model(table string) string {
	modelSlice := strings.Split(table, "_")
	for i, s := range modelSlice {
		modelSlice[i] = estrings.FirstUpper(s)
	}
	return strings.Join(modelSlice, "")
}
