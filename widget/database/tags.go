package database

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/whf-sky/efficient/widget/dtype"
	"github.com/whf-sky/efficient/widget/tools/numeric"
	"github.com/whf-sky/efficient/widget/validation"
	"math"
	"reflect"
	"strconv"
	"strings"
)

type dirvers map[string]dbnames

type dbnames map[string]tables

type tables map[string]columns

type columns map[string]Tags

var dirs = dirvers{}

type TagsInterface interface {
	Get(key string)
	Set(key string, val interface{})
}

func NewTags() *Tags {
	return &Tags{}
}

type Tags struct {
	Column         string       //字段
	Comment        string       //注释
	Type           string       //类型
	Unsigned       bool         //无符号
	Precision      int64        //数字精度
	Scale          int64        //小数位精度
	Length         int64        //字符串字符长度
	OctetLength    int64        //字符串字节长度
	AutoIncrement  bool         //自增
	PrimaryKey     bool         //主键
	Unique         bool         //唯一索引
	Default        interface{}  //默认值
	Validators     validation.V //检查字段内容
	Null           bool         //不能为空
	AutoCreateTime bool         //自动创建时间
	AutoUpdateTime bool         //自动跟新时间
	Enum           []string     //枚举
}

//parseModel 对model进行分析，获取标签信息
func parseModel(model ModelInterface) {
	if _, ok := dirs[model.DriverName()]; !ok {
		dirs[model.DriverName()] = dbnames{}
	}
	if _, ok := dirs[model.DriverName()][model.Dbname()]; !ok {
		dirs[model.DriverName()][model.Dbname()] = tables{}
	}
	if _, ok := dirs[model.DriverName()][model.Dbname()][model.Key()]; !ok {
		dirs[model.DriverName()][model.Dbname()][model.Key()] = columns{}
	}
	elem := reflect.TypeOf(model).Elem()
	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		column := conv(field.Tag, "column", "string", "").(string)
		if column == "" {
			continue
		}
		tags := Tags{
			Column:         column,
			Type:           conv(field.Tag, "type", "string", "").(string),
			Enum:           conv(field.Tag, "enum", "[]string", []string{}).([]string),
			AutoIncrement:  conv(field.Tag, "auto_increment", "bool", false).(bool),
			PrimaryKey:     conv(field.Tag, "primary_key", "bool", false).(bool),
			Unique:         conv(field.Tag, "unique", "bool", false).(bool),
			Unsigned:       conv(field.Tag, "unsigned", "bool", false).(bool),
			Precision:      conv(field.Tag, "precision", "int64", int64(0)).(int64),
			Scale:          conv(field.Tag, "scale", "int64", int64(0)).(int64),
			Length:         conv(field.Tag, "length", "int64", int64(0)).(int64),
			OctetLength:    conv(field.Tag, "octet_length", "int64", int64(0)).(int64),
			Null:           conv(field.Tag, "null", "bool", false).(bool),
			Default:        conv(field.Tag, "default", "string", "").(string),
			Comment:        conv(field.Tag, "comment", "string", "").(string),
			AutoCreateTime: conv(field.Tag, "on_insert_time", "bool", false).(bool),
			AutoUpdateTime: conv(field.Tag, "on_update_time", "bool", false).(bool),
			Validators:     validators(field.Tag, conv(field.Tag, "validators", "string", "").(string)),
		}
		tags.Validators = fieldValidators(tags)
		tags.Default = formattingDefault(tags)
		dirs[model.DriverName()][model.Dbname()][model.Key()][column] = tags
	}
}

//formattingDefault 格式化默认值
func formattingDefault(tags Tags) interface{} {
	s := tags.Default.(string)
	if s == "" {
		return nil
	}
	switch tags.Type {
	case "time", "year", "char", "varchar", "tinytext", "text", "mediumtext", "longtext", "enum":
		return s
	case "bit":
		if tags.Precision == 1 {
			if s[2:] == "1" {
				return 1
			}
			return 0
		}
		return numeric.Bin2dec(s[2:])
	case "set":
		var set dtype.Set
		_ = set.Scan([]byte(s))
		return set
	case "date":
		var date dtype.Date
		_ = date.Scan([]byte(s))
		return date
	case "datetime", "timestamp":
		var datetime dtype.Time
		_ = datetime.Scan([]byte(s))
		return datetime
	case "tinyint":
		if tags.Unsigned {
			var tinyint dtype.NullUint8
			_ = tinyint.Scan([]byte(s))
			return tinyint
		}
		var tinyint dtype.NullInt8
		_ = tinyint.Scan([]byte(s))
		return tinyint
	case "smallint":
		if tags.Unsigned {
			var smallint dtype.NullUint16
			_ = smallint.Scan([]byte(s))
			return smallint
		}
		var smallint sql.NullInt16
		_ = smallint.Scan([]byte(s))
		return smallint
	case "mediumint", "int", "integer":
		if tags.Unsigned {
			var num dtype.NullUint32
			_ = num.Scan([]byte(s))
			return num
		}
		var num sql.NullInt32
		_ = num.Scan([]byte(s))
		return num
	case "bigint":
		if tags.Unsigned {
			var num dtype.NullUint64
			_ = num.Scan([]byte(s))
			return num
		}
		var num sql.NullInt64
		_ = num.Scan([]byte(s))
		return num
	case "decimal", "numeric":
		var num decimal.Decimal
		_ = num.Scan([]byte(s))
		return num
	case "double", "float", "real":
		var num sql.NullFloat64
		_ = num.Scan([]byte(s))
		return num
	}
	return nil
}

//numValidators 字段的数字类型验证器
func numValidators(tags Tags, bit float64) validation.V {
	if tags.Unsigned {
		switch bit {
		case 8:
			tags.Validators["gte"] = uint8(0)
			tags.Validators["lte"] = uint8(math.MaxUint8)
		case 16:
			tags.Validators["gte"] = uint16(0)
			tags.Validators["lte"] = uint16(math.MaxUint16)
		case 24, 32:
			tags.Validators["gte"] = uint32(0)
			tags.Validators["lte"] = uint32(math.MaxUint32)
		case 64:
			tags.Validators["gte"] = 0
			tags.Validators["lte"] = uint64(math.MaxUint64)
		}
		return tags.Validators
	}
	switch bit {
	case 8:
		tags.Validators["lte"] = int8(math.MaxInt8)
		tags.Validators["gte"] = int8(math.MinInt8)
	case 16:
		tags.Validators["lte"] = int16(math.MaxInt16)
		tags.Validators["gte"] = int16(math.MinInt16)
	case 24, 32:
		tags.Validators["lte"] = int32(math.MaxInt32)
		tags.Validators["gte"] = int32(math.MinInt32)
	case 64:
		tags.Validators["lte"] = int64(math.MaxInt64)
		tags.Validators["gte"] = int64(math.MinInt64)
	}
	return tags.Validators
}

//fieldValidators 为字段添加数据验证器
func fieldValidators(tags Tags) validation.V {
	switch tags.Type {
	//字符串
	case "char", "varchar", "tinytext", "text", "mediumtext", "longtext":
		tags.Validators["len-lte"] = int(tags.Length)
	//枚举
	case "enum": //单选
		tags.Validators["in"] = tags.Enum
	case "set": //多选
		tags.Validators["in-multi"] = tags.Enum
	case "tinyint": //8
		tags.Validators = numValidators(tags, 8)
	case "smallint": //16
		tags.Validators = numValidators(tags, 16)
	case "mediumint": //24
		tags.Validators = numValidators(tags, 24)
	case "int", "integer": //32
		tags.Validators = numValidators(tags, 32)
	case "bigint": //64
		tags.Validators = numValidators(tags, 64)
	}
	return tags.Validators
}

//conv 类型数据转换
func conv(tag reflect.StructTag, key string, gtype string, defaultVal interface{}) interface{} {
	val, ok := tag.Lookup(key)
	if !ok {
		return defaultVal
	}
	switch gtype {
	case "bool":
		b, _ := strconv.ParseBool(val)
		return b
	case "string":
		return val
	case "[]string":
		//v-in="[]string:{'1','2'}"
		return strings.Split(val[2:len(val)-2], "','")
	case "int64":
		num, _ := strconv.ParseInt(val, 10, 64)
		return num
	}
	return defaultVal
}

//validators tag中的自定义验证器
//for example:
//`validators="lte,gte" v-lte="10" v-gte="5"`
func validators(tag reflect.StructTag, s string) validation.V {
	V := validation.V{}
	if len(s) == 0 {
		return V
	}
	validators := strings.Split(s, ",") //验证器
	for _, validator := range validators {
		//获取验证器参数
		val, ok := tag.Lookup("v-" + validator)
		if !ok {
			V[validator] = nil
			continue
		}

		//验证器参数类型数据，:前为参数数据类型，:后为参数数据
		typeParam := strings.SplitN(val, ":", 2)
		switch typeParam[0] {
		case "int8":
			num, err := strconv.Atoi(typeParam[1])
			if err != nil {
				V[validator] = int8(0)
			} else {
				V[validator] = int8(num)
			}
		case "uint8":
			num, err := strconv.Atoi(typeParam[1])
			if err != nil {
				V[validator] = uint8(0)
			} else {
				V[validator] = uint8(num)
			}
		case "int16":
			num, err := strconv.Atoi(typeParam[1])
			if err != nil {
				V[validator] = int16(0)
			} else {
				V[validator] = int16(num)
			}
		case "uint16":
			num, err := strconv.Atoi(typeParam[1])
			if err != nil {
				V[validator] = uint16(0)
			} else {
				V[validator] = uint16(num)
			}
		case "int32":
			num, err := strconv.Atoi(typeParam[1])
			if err != nil {
				V[validator] = int32(0)
			} else {
				V[validator] = int32(num)
			}
		case "uint32":
			num, err := strconv.Atoi(typeParam[1])
			if err != nil {
				V[validator] = uint32(0)
			} else {
				V[validator] = uint32(num)
			}
		case "int":
			num, err := strconv.Atoi(typeParam[1])
			if err != nil {
				V[validator] = 0
			} else {
				V[validator] = num
			}
		case "uint":
			num, err := strconv.Atoi(typeParam[1])
			if err != nil {
				V[validator] = uint(0)
			} else {
				V[validator] = uint(num)
			}
		case "int64":
			num, err := strconv.ParseInt(typeParam[1], 10, 64)
			if err != nil {
				V[validator] = int64(0)
			} else {
				V[validator] = num
			}
		case "uint64":
			num, err := strconv.ParseInt(typeParam[1], 10, 64)
			if err != nil {
				V[validator] = int64(0)
			} else {
				V[validator] = num
			}
		case "string":
			V[validator] = typeParam[1]
		case "[]string":
			s := typeParam[1]
			V[validator] = strings.Split(s[2:len(s)-2], "','")
		case "float32":
			num, err := strconv.ParseFloat(typeParam[1], 64)
			if err != nil {
				V[validator] = float32(0)
			} else {
				V[validator] = float32(num)
			}
		case "float64":
			num, err := strconv.ParseFloat(typeParam[1], 64)
			if err != nil {
				V[validator] = float64(0)
			} else {
				V[validator] = float64(num)
			}
		case "bool":
			b, err := strconv.ParseBool(typeParam[1])
			if err != nil {
				V[validator] = false
			} else {
				V[validator] = b
			}
		default:
		}
	}
	return V
}
