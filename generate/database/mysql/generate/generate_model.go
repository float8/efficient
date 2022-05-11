package generate

import (
	"github.com/whf-sky/efficient/database"
	"github.com/whf-sky/efficient/generate/database/mysql/model"
	estrings "github.com/whf-sky/efficient/tools/strings"
	"strings"
)

func newGenerateModel(tableName string, columns []database.ModelInterface, pkg string) *generateModel {
	return &generateModel{
		tableName: tableName,
		columns:   columns,
		pkg:       pkg,
		imports: []string{
			"encoding/json",
		},
		sections: map[string]string{
			"package":   "package #package#",
			"import":    "import (\n\t\"#imports#\"\n)",
			"new":       "func New#model#() *#model# {\n\tm := &#model#{}\n\tm.init(\"#tablename#\")\n\treturn m\n}",
			"ptrs":      "func (#modelvar# *#model#) Ptrs() map[string]interface{} {\n\treturn map[string]interface{}{#ptrs#}\n}",
			"struct":    "type #model# struct {\n\tModel\n\t#fields#\n}",
			"event":     "func (#modelvar# *#model#) InsertEvent() {}\n\nfunc (#modelvar# *#model#) UpdateEvent() {}",
			"tableName": "func (#modelvar# *#model#) TableName() string {\n\treturn \"#tablename#\"\n}",
			"getFields": "func (#modelvar# *#model#) Get#field#() #type# {\n\treturn #modelvar#.#field#\n}",
			"setFields": "func (#modelvar# *#model#) Set#field#(value #type#) *#model# {\n\t#modelvar#.AddAssColumns(\"#column#\")\n\t#modelvar#.#field# = value\n\treturn #modelvar#\n}",
			"get":       "func (#modelvar# *#model#) Get(key string) interface{} {\n\tswitch key {\n\t#cases#\n\t}\n\treturn nil\n}",
			"set":       "func (#modelvar# *#model#) Set(key string, val interface{}) {\n\tswitch key {\n\t#cases#\n\t}\n}",
			"toString":  "func (#modelvar# *#model#) ToString() string {\n\tbytes, _ := json.Marshal(#modelvar#)\n\treturn string(bytes)\n}",
		},
		sectionsOrder:     []string{"package", "import", "new", "struct", "tableName", "ptrs", "get", "set", "getFields", "setFields", "event", "toString"},
		groups:            map[string][]string{},
		repetitionImports: map[string]bool{},
		ptrs:              [][]string{},
	}
}

type field struct {
	Var  string
	Type string
	Tag  string
}

type generateModel struct {
	pkg               string
	tableName         string
	columns           []database.ModelInterface
	modelVar          string
	repetitionImports map[string]bool
	imports           []string
	sections          map[string]string
	sectionsOrder     []string
	groups            map[string][]string
	structFields      []field
	ptrs              [][]string
	maxFieldLen       int
}

func (g *generateModel) generate() string {
	g.packageHandle()
	g.tableHandle(g.tableName)
	g.columnsHandle()
	code := []string{}
	for _, key := range g.sectionsOrder {
		code = append(code, g.sections[key])
	}
	return strings.Join(code, "\n\n")
}

func (g *generateModel) packageHandle() {
	g.sections["package"] = strings.ReplaceAll(g.sections["package"], "#package#", g.pkg)
}

func (g *generateModel) tableHandle(table string) {
	table = strings.ToLower(table)
	modelSlice := strings.Split(table, "_")
	for i, s := range modelSlice {
		modelSlice[i] = estrings.FirstUpper(s)
	}

	g.modelVar = strings.ToLower(string(modelSlice[0][0]))

	for key, ele := range g.sections {
		ele = strings.ReplaceAll(ele, "#tablename#", table)
		ele = strings.ReplaceAll(ele, "#drivername#", "mysql")
		ele = strings.ReplaceAll(ele, "#model#", strings.Join(modelSlice, ""))
		ele = strings.ReplaceAll(ele, "#modelvar#", g.modelVar)
		g.sections[key] = ele
	}
}

func (g *generateModel) ptrsHandle() {
	ptrs := []string{}
	for _, ptr := range g.ptrs {
		plen := len(ptr[0])
		if plen < g.maxFieldLen {
			ptrs = append(ptrs, "\""+ptr[0]+"\":"+strings.Repeat(" ", g.maxFieldLen-plen+1)+ptr[1])
			continue
		} else {
			ptrs = append(ptrs, "\""+ptr[0]+"\": "+ptr[1])
		}

	}
	g.sections["ptrs"] = strings.ReplaceAll(g.sections["ptrs"], "#ptrs#", "\n\t\t"+strings.Join(ptrs, ",\n\t\t")+",\n\t")
}

func (g *generateModel) columnsHandle() {
	for _, column := range g.columns {
		cln := column.(*model.Columns)
		g.structFieldHandle(cln)
	}
	g.importHandle()
	g.structHandle()
	g.ptrsHandle()
	g.getHandle()
	g.setHandle()
	g.getFieldsHandle()
	g.setFieldsHandle()
}

func (g *generateModel) setFieldsHandle() {
	g.sections["setFields"] = strings.Join(g.groups["setFields"], "\n\n")
}

func (g *generateModel) getFieldsHandle() {
	g.sections["getFields"] = strings.Join(g.groups["getFields"], "\n\n")
}

func (g *generateModel) structFieldHandle(column *model.Columns) {
	tags, tagsStr := g.tags(column)
	goType := g.goType(tags)
	fieldStr := g.field(tags["column"])

	g.ptrs = append(g.ptrs, []string{
		tags["column"],
		"&" + g.modelVar + "." + fieldStr,
	})

	clnLen := len(tags["column"])
	if clnLen > g.maxFieldLen {
		g.maxFieldLen = clnLen
	}

	g.structFields = append(g.structFields, field{
		Var:  fieldStr,
		Type: goType,
		Tag:  tagsStr,
	})
	g.getFieldHandle(fieldStr, goType)
	g.setFieldHandle(tags["column"], fieldStr, goType)
	g.getCases(tags["column"], fieldStr)
	g.setCases(tags["column"], fieldStr, goType)
}

func (g *generateModel) goType(tags map[string]string) string {
	mtype := tags["type"]
	if mtype == "bit" && tags["precision"] != "1" {
		mtype = "bits"
	}

	gType := goType{}
	if tags["null"] == "false" {
		gType = types[mtype]
	} else {
		gType = nullTypes[mtype]
	}

	if tags["unsigned"] == "true" && gType.unsigned.name != "" {
		gUtype := gType.unsigned
		_, ok := g.repetitionImports[gUtype.pkg]
		if gUtype.pkg != "" && !ok {
			g.repetitionImports[gUtype.pkg] = true
			g.imports = append(g.imports, gUtype.pkg)
		}
		return gUtype.name
	}
	_, ok := g.repetitionImports[gType.pkg]
	if gType.pkg != "" && !ok {
		g.repetitionImports[gType.pkg] = true
		g.imports = append(g.imports, gType.pkg)
	}
	return gType.name

}

func (g *generateModel) mysqlConvGoType(typename string, tags map[string]string) string {
	tp := ""
	if tags["unsigned"] == "true" && strings.Index(tags["type"], "int") > -1 {
		tp += "u" + typename
	} else {
		tp += typename
	}
	return tp
}

func (g *generateModel) tags(column *model.Columns) (tags map[string]string, tagsStr string) {
	order := []string{"column", "type", "auto_increment", "primary_key",
		"unique", "unsigned", "precision", "scale", "length", "octet_length", "null", "enum",
		"on_insert_time", "on_update_time", "default", "comment"}

	tags = map[string]string{
		"column":       column.GetColumnName(),
		"type":         column.GetDataType(),
		"comment":      column.GetColumnComment(),
		"precision":    column.GetNumericPrecision(),
		"scale":        column.GetNumericScale(),
		"length":       column.GetCharacterMaximumLength(),
		"octet_length": column.GetCharacterOctetLength(),

		"primary_key": func(s string) string {
			if s == "PRI" {
				return "true"
			}
			return "false"
		}(column.GetColumnKey()),

		"unique": func(s string) string {
			if s == "UNI" {
				return "true"
			}
			return "false"
		}(column.GetColumnKey()),

		"null": func(s string) string {
			if s == "YES" {
				return "true"
			}
			return "false"
		}(column.GetIsNullable()),

		"default": func(s string) string {
			s = strings.Trim(s, "'")
			if s == "current_timestamp()" || s == "NULL" || s == "" {
				return ""
			}
			return s
		}(column.GetColumnDefault()),

		"unsigned": func(ctype string) string {
			if strings.Index(ctype, "unsigned") > -1 {
				return "true"
			}
			return "false"
		}(column.GetColumnType()),

		"auto_increment": func(extra string) string {
			if extra == "auto_increment" {
				return "true"
			}
			return "false"
		}(column.GetExtra()),

		"on_update_time": func(s string) string {
			if s == "on update current_timestamp()" {
				return "true"
			}
			return "false"
		}(column.GetExtra()),

		"on_insert_time": func(s string) string {
			if s == "current_timestamp()" {
				return "true"
			}
			return "false"
		}(column.GetColumnDefault()),

		"enum": func(dType, s string) string {
			if s != "" && dType == "enum" {
				r := []rune(s)
				return "{" + string(r[5:len(r)-1]) + "}"
			}
			if s != "" && dType == "set" {
				r := []rune(s)
				return "{" + string(r[4:len(r)-1]) + "}"
			}
			return ""
		}(column.GetDataType(), column.GetColumnType()),
	}

	tagsSlice := []string{}
	for _, key := range order {
		if tags[key] == "false" {
			continue
		}
		if tags[key] != "" {
			tagsSlice = append(tagsSlice, key+":\""+tags[key]+"\"")
		}
	}

	tagsStr = "`" + strings.Join(tagsSlice, " ") + "`"

	return
}

func (g *generateModel) field(s string) string {
	fieldStr := ""
	nameArr := strings.Split(s, "_")
	for _, name := range nameArr {
		fieldStr += estrings.FirstUpper(name)
	}
	return fieldStr
}

func (g *generateModel) getFieldHandle(field, goType string) {
	getField := g.sections["getFields"]
	getField = strings.ReplaceAll(getField, "#field#", field)
	getField = strings.ReplaceAll(getField, "#type#", goType)
	g.groups["getFields"] = append(g.groups["getFields"], getField)
}

func (g *generateModel) setFieldHandle(column, field, goType string) {
	setField := g.sections["setFields"]
	setField = strings.ReplaceAll(setField, "#field#", field)
	setField = strings.ReplaceAll(setField, "#type#", goType)
	setField = strings.ReplaceAll(setField, "#column#", column)
	g.groups["setFields"] = append(g.groups["setFields"], setField)
}

func (g *generateModel) getCases(column, field string) {
	//tpl := "case \"#column#\":\n\t\treturn #modelvar#.Get#field#()"
	tpl := "case \"#column#\":\n\t\treturn #modelvar#.#field#"
	caseStr := strings.ReplaceAll(tpl, "#modelvar#", g.modelVar)
	caseStr = strings.ReplaceAll(caseStr, "#column#", column)
	caseStr = strings.ReplaceAll(caseStr, "#field#", field)

	g.groups["getCases"] = append(g.groups["getCases"], caseStr)
}

func (g *generateModel) setCases(column, field, goType string) {
	//tpl := "case \"#column#\":\n\t\t#modelvar#.Set#field#(val.(#type#))"
	tpl := "case \"#column#\":\n\t\t#modelvar#.#field# = val.(#type#)"
	caseStr := strings.ReplaceAll(tpl, "#modelvar#", g.modelVar)
	caseStr = strings.ReplaceAll(caseStr, "#column#", column)
	caseStr = strings.ReplaceAll(caseStr, "#field#", field)
	caseStr = strings.ReplaceAll(caseStr, "#type#", goType)
	g.groups["setCases"] = append(g.groups["setCases"], caseStr)
}

func (g *generateModel) structHandle() {
	fields := []string{}
	varMaxLen := 0
	typeMaxLen := 0
	for _, structField := range g.structFields {
		vl := len(structField.Var)
		if vl > varMaxLen {
			varMaxLen = vl
		}
		tl := len(structField.Type)
		if tl > typeMaxLen {
			typeMaxLen = tl
		}
	}

	for _, structField := range g.structFields {
		vl := len(structField.Var)
		tl := len(structField.Type)
		stype := structField.Type + strings.Repeat(" ", typeMaxLen-tl+1)
		svar := structField.Var + strings.Repeat(" ", varMaxLen-vl+1)
		fields = append(fields, svar+stype+structField.Tag)
	}

	fieldsStr := strings.Join(fields, "\n\t")
	g.sections["struct"] = strings.ReplaceAll(g.sections["struct"], "#fields#", fieldsStr)
}

func (g *generateModel) getHandle() {
	fields := strings.Join(g.groups["getCases"], "\n\t")
	g.sections["get"] = strings.ReplaceAll(g.sections["get"], "#cases#", fields)
}

func (g *generateModel) setHandle() {
	fields := strings.Join(g.groups["setCases"], "\n\t")
	g.sections["set"] = strings.ReplaceAll(g.sections["set"], "#cases#", fields)
}

func (g *generateModel) importHandle() {

	imports := strings.Join(g.imports, "\"\n\t\"")
	g.sections["import"] = strings.ReplaceAll(g.sections["import"], "#imports#", imports)
}
