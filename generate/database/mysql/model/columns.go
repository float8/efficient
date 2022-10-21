package model

import (
	"database/sql"
	"github.com/float8/efficient/database"
)

func NewColumns() *Columns {
	users := &Columns{}
	users.Init("mysql", "information_schema", "_columns")
	return users
}

type Columns struct {
	database.Model
	TableCatalog    sql.NullString `column:"TABLE_CATALOG"`
	TableSchema     sql.NullString `column:"TABLE_SCHEMA"`     //数据库名称
	TName           sql.NullString `column:"TABLE_NAME"`       //表名
	ColumnName      sql.NullString `column:"COLUMN_NAME"`      //字段名
	OrdinalPosition sql.NullString `column:"ORDINAL_POSITION"` //字段编号
	ColumnDefault   sql.NullString `column:"COLUMN_DEFAULT"`   //字段默认值。
	IsNullable      sql.NullString `column:"IS_NULLABLE"`      //字段是否可以是NULL。 该列记录的值是YES或者NO。
	DataType        sql.NullString `column:"DATA_TYPE"`        //数据类型。 里面的值是字符串，比如varchar，float，int。

	//字段的最大字符数。
	// 假如字段设置为varchar(50)，那么这一列记录的值就是50。
	// 该列只适用于二进制数据，字符，文本，图像数据。其他类型数据比如int，float，datetime等，在该列显示为NULL。
	CharacterMaximumLength sql.NullString `column:"CHARACTER_MAXIMUM_LENGTH"`

	//字段的最大字节数。
	//和最大字符数一样，只适用于二进制数据，字符，文本，图像数据，其他类型显示为NULL。
	//和最大字符数的数值有比例关系，和字符集有关。比如UTF8类型的表，最大字节数就是最大字符数的3倍。
	CharacterOctetLength sql.NullString `column:"CHARACTER_OCTET_LENGTH"`

	//数字精度。
	// 适用于各种数字类型比如int，float之类的。
	//如果字段设置为int(10)，那么在该列保存的数值是9，少一位，还没有研究原因。
	//如果字段设置为float(10,3)，那么在该列报错的数值是10。
	//非数字类型显示为在该列NULL。
	NumericPrecision sql.NullString `column:"NUMERIC_PRECISION"`

	//小数位数。
	// 和数字精度一样，适用于各种数字类型比如int，float之类。
	//如果字段设置为int(10)，那么在该列保存的数值是0，代表没有小数。
	//如果字段设置为float(10,3)，那么在该列报错的数值是3。
	//非数字类型显示为在该列NULL。
	NumericScale sql.NullString `column:"NUMERIC_SCALE"`

	//datetime类型和SQL-92interval类型数据库的子类型代码。
	// 我本地datetime类型的字段在该列显示为0。
	//其他类型显示为NULL。
	DatatimePrecision sql.NullString `column:"DATETIME_PRECISION"`

	CharacterSetName sql.NullString `column:"CHARACTER_SET_NAME"` //字段字符集名称。比如utf8。

	//字符集排序规则。
	// 比如utf8_general_ci，是不区分大小写一种排序规则。
	//utf8_general_cs，是区分大小写的排序规则。
	CollationName sql.NullString `column:"COLLATION_NAME"`

	ColumnType           sql.NullString `column:"COLUMN_TYPE"`    //字段类型。比如float(9,3)，varchar(50)。
	ColumnKey            sql.NullString `column:"COLUMN_KEY"`     //索引类型。可包含的值有PRI，代表主键，UNI，代表唯一键，MUL，可重复。
	Extra                sql.NullString `column:"EXTRA"`          //其他信息。比如主键的auto_increment。
	Privileges           sql.NullString `column:"PRIVILEGES"`     //权限.多个权限用逗号隔开，比如 select,insert,update,references
	ColumnComment        sql.NullString `column:"COLUMN_COMMENT"` //字段注释
	IsGenerated          sql.NullString `column:"IS_GENERATED"`
	GenerationExpression sql.NullString `column:"GENERATION_EXPRESSION"` //组合字段的公式。
}

func (c *Columns) TableName() string {
	return "columns"
}

func (c *Columns) Ptrs() map[string]interface{} {
	return map[string]interface{}{
		"TABLE_CATALOG":            &c.TableCatalog,
		"TABLE_SCHEMA":             &c.TableSchema,
		"TABLE_NAME":               &c.TName,
		"COLUMN_NAME":              &c.ColumnName,
		"ORDINAL_POSITION":         &c.OrdinalPosition,
		"COLUMN_DEFAULT":           &c.ColumnDefault,
		"IS_NULLABLE":              &c.IsNullable,
		"DATA_TYPE":                &c.DataType,
		"CHARACTER_MAXIMUM_LENGTH": &c.CharacterMaximumLength,
		"CHARACTER_OCTET_LENGTH":   &c.CharacterOctetLength,
		"NUMERIC_PRECISION":        &c.NumericPrecision,
		"NUMERIC_SCALE":            &c.NumericScale,
		"DATETIME_PRECISION":       &c.DatatimePrecision,
		"CHARACTER_SET_NAME":       &c.CharacterSetName,
		"COLLATION_NAME":           &c.CollationName,
		"COLUMN_TYPE":              &c.ColumnType,
		"COLUMN_KEY":               &c.ColumnKey,
		"EXTRA":                    &c.Extra,
		"PRIVILEGES":               &c.Privileges,
		"COLUMN_COMMENT":           &c.ColumnComment,
		"IS_GENERATED":             &c.IsGenerated,
		"GENERATION_EXPRESSION":    &c.GenerationExpression,
	}
}

func (c *Columns) Get(key string) interface{} {
	switch key {
	case "TABLE_CATALOG":
		return c.GetTableCatalog()
	case "TABLE_SCHEMA":
		return c.GetTableSchema()
	case "TABLE_NAME":
		return c.GetTName()
	case "COLUMN_NAME":
		return c.GetColumnName()
	case "COLUMN_DEFAULT":
		return c.GetColumnDefault()
	case "IS_NULLABLE":
		return c.GetIsNullable()
	case "DATA_TYPE":
		return c.GetDataType()
	case "CHARACTER_SET_NAME":
		return c.GetCharacterSetName()
	case "COLLATION_NAME":
		return c.GetCollationName()
	case "COLUMN_TYPE":
		return c.GetColumnType()
	case "COLUMN_KEY":
		return c.GetColumnKey()
	case "EXTRA":
		return c.GetExtra()
	case "PRIVILEGES":
		return c.GetPrivileges()
	case "COLUMN_COMMENT":
		return c.GetColumnComment()
	case "IS_GENERATED":
		return c.GetIsGenerated()
	case "GENERATION_EXPRESSION":
		return c.GetGenerationExpression()

	case "CHARACTER_MAXIMUM_LENGTH":
		return c.GetCharacterMaximumLength()
	case "CHARACTER_OCTET_LENGTH":
		return c.GetCharacterOctetLength()
	case "NUMERIC_PRECISION":
		return c.GetNumericPrecision()
	case "NUMERIC_SCALE":
		return c.GetNumericScale()
	case "DATETIME_PRECISION":
		return c.GetDatatimePrecision()
	case "ORDINAL_POSITION":
		return c.GetOrdinalPosition()
	}
	return nil
}

func (c *Columns) Set(key string, val interface{}) {
	v := val.(string)
	switch key {
	case "TABLE_CATALOG":
		c.SetTableCatalog(v)
	case "TABLE_SCHEMA":
		c.SetTableSchema(v)
	case "TABLE_NAME":
		c.SetTName(v)
	case "COLUMN_NAME":
		c.SetColumnName(v)
	case "COLUMN_DEFAULT":
		c.SetColumnDefault(v)
	case "IS_NULLABLE":
		c.SetIsNullable(v)
	case "DATA_TYPE":
		c.SetDataType(v)
	case "CHARACTER_SET_NAME":
		c.SetCharacterSetName(v)
	case "COLLATION_NAME":
		c.SetCollationName(v)
	case "COLUMN_TYPE":
		c.SetColumnType(v)
	case "COLUMN_KEY":
		c.SetColumnKey(v)
	case "EXTRA":
		c.SetExtra(v)
	case "PRIVILEGES":
		c.SetPrivileges(v)
	case "COLUMN_COMMENT":
		c.SetColumnComment(v)
	case "IS_GENERATED":
		c.SetIsGenerated(v)
	case "GENERATION_EXPRESSION":
		c.SetGenerationExpression(v)

	case "CHARACTER_MAXIMUM_LENGTH":
		c.SetCharacterMaximumLength(v)
	case "CHARACTER_OCTET_LENGTH":
		c.SetCharacterOctetLength(v)
	case "NUMERIC_PRECISION":
		c.SetNumericPrecision(v)
	case "NUMERIC_SCALE":
		c.SetNumericScale(v)
	case "DATETIME_PRECISION":
		c.SetDatatimePrecision(v)
	case "ORDINAL_POSITION":
		c.SetOrdinalPosition(v)
	}
}

func (c *Columns) GetTableCatalog() string {
	return c.TableCatalog.String
}

func (c *Columns) GetTableSchema() string {
	return c.TableSchema.String
}

func (c *Columns) GetTName() string {
	return c.TName.String
}

func (c *Columns) GetColumnName() string {
	return c.ColumnName.String
}

func (c *Columns) GetOrdinalPosition() string {
	return c.OrdinalPosition.String
}

func (c *Columns) GetColumnDefault() string {
	return c.ColumnDefault.String
}

func (c *Columns) GetIsNullable() string {
	return c.IsNullable.String
}

func (c *Columns) GetDataType() string {
	return c.DataType.String
}

func (c *Columns) GetCharacterMaximumLength() string {
	return c.CharacterMaximumLength.String
}

func (c *Columns) GetCharacterOctetLength() string {
	return c.CharacterOctetLength.String
}

func (c *Columns) GetNumericPrecision() string {
	return c.NumericPrecision.String
}

func (c *Columns) GetNumericScale() string {
	return c.NumericScale.String
}

func (c *Columns) GetDatatimePrecision() string {
	return c.DatatimePrecision.String
}

func (c *Columns) GetCharacterSetName() string {
	return c.CharacterSetName.String
}

func (c *Columns) GetCollationName() string {
	return c.CollationName.String
}

func (c *Columns) GetColumnType() string {
	return c.ColumnType.String
}

func (c *Columns) GetColumnKey() string {
	return c.ColumnKey.String
}

func (c *Columns) GetExtra() string {
	return c.Extra.String
}

func (c *Columns) GetPrivileges() string {
	return c.Privileges.String
}

func (c *Columns) GetColumnComment() string {
	return c.ColumnComment.String
}

func (c *Columns) GetIsGenerated() string {
	return c.IsGenerated.String
}

func (c *Columns) GetGenerationExpression() string {
	return c.TableCatalog.String
}

//---------------------set------------------------------

func (c *Columns) SetTableCatalog(v string) *Columns {
	c.TableCatalog = sql.NullString{String: v}
	return c
}

func (c *Columns) SetTableSchema(v string) *Columns {
	c.TableSchema = sql.NullString{String: v}
	return c
}

func (c *Columns) SetTName(v string) *Columns {
	c.TName = sql.NullString{String: v}
	return c
}

func (c *Columns) SetColumnName(v string) *Columns {
	c.ColumnName = sql.NullString{String: v}
	return c
}

func (c *Columns) SetOrdinalPosition(v string) *Columns {
	c.OrdinalPosition = sql.NullString{String: v}
	return c
}

func (c *Columns) SetColumnDefault(v string) *Columns {
	c.ColumnDefault = sql.NullString{String: v}
	return c
}

func (c *Columns) SetIsNullable(v string) *Columns {
	c.IsNullable = sql.NullString{String: v}
	return c
}

func (c *Columns) SetDataType(v string) *Columns {
	c.DataType = sql.NullString{String: v}
	return c
}

func (c *Columns) SetCharacterMaximumLength(v string) *Columns {
	c.CharacterMaximumLength = sql.NullString{String: v}
	return c
}

func (c *Columns) SetCharacterOctetLength(v string) *Columns {
	c.CharacterOctetLength = sql.NullString{String: v}
	return c
}

func (c *Columns) SetNumericPrecision(v string) *Columns {
	c.NumericPrecision = sql.NullString{String: v}
	return c
}

func (c *Columns) SetNumericScale(v string) *Columns {
	c.NumericScale = sql.NullString{String: v}
	return c
}

func (c *Columns) SetDatatimePrecision(v string) *Columns {
	c.DatatimePrecision = sql.NullString{String: v}
	return c
}

func (c *Columns) SetCharacterSetName(v string) *Columns {
	c.CharacterSetName = sql.NullString{String: v}
	return c
}

func (c *Columns) SetCollationName(v string) *Columns {
	c.CollationName = sql.NullString{String: v}
	return c
}

func (c *Columns) SetColumnType(v string) *Columns {
	c.ColumnType = sql.NullString{String: v}
	return c
}

func (c *Columns) SetColumnKey(v string) *Columns {
	c.ColumnKey = sql.NullString{String: v}
	return c
}

func (c *Columns) SetExtra(v string) *Columns {
	c.Extra = sql.NullString{String: v}
	return c
}

func (c *Columns) SetPrivileges(v string) *Columns {
	c.Privileges = sql.NullString{String: v}
	return c
}

func (c *Columns) SetColumnComment(v string) *Columns {
	c.ColumnComment = sql.NullString{String: v}
	return c
}

func (c *Columns) SetIsGenerated(v string) *Columns {
	c.IsGenerated = sql.NullString{String: v}
	return c
}

func (c *Columns) SetGenerationExpression(v string) *Columns {
	c.TableCatalog = sql.NullString{String: v}
	return c
}
