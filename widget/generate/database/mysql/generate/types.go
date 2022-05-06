package generate

type goType struct {
	name     string
	pkg      string
	unsigned goUtype
}

type goUtype struct {
	name string
	pkg  string
}

func RegisterType(unsigned bool, mType, gType string, pkgs ...string) {
	if unsigned {
		registerUtype(mType, gType, pkgs...)
		return
	}
	if len(pkgs) > 0 {
		types[mType] = goType{name: gType, pkg: pkgs[0]}
		return
	}
	types[mType] = goType{name: gType}
}

func registerUtype(mType, gType string, pkgs ...string) {
	if len(pkgs) > 0 {
		types[mType] = goType{
			name: types[mType].name,
			pkg:  types[mType].pkg,
			unsigned: goUtype{
				name: gType,
				pkg:  pkgs[0],
			},
		}
		return
	}
	types[mType] = goType{
		name: types[mType].name,
		pkg:  types[mType].pkg,
		unsigned: goUtype{
			name: gType,
		},
	}
}

func RegisterNullType(unsigned bool, mType, gType string, pkgs ...string) {
	if unsigned {
		registerNullUtype(mType, gType, pkgs...)
		return
	}
	if len(pkgs) > 0 {
		nullTypes[mType] = goType{name: gType, pkg: pkgs[0]}
		return
	}
	types[mType] = goType{name: gType}
}

func registerNullUtype(mType, gType string, pkgs ...string) {
	if len(pkgs) > 0 {
		nullTypes[mType] = goType{
			name: nullTypes[mType].name,
			pkg:  nullTypes[mType].pkg,
			unsigned: goUtype{
				name: gType,
				pkg:  pkgs[0],
			},
		}
		return
	}
	nullTypes[mType] = goType{
		name: nullTypes[mType].name,
		pkg:  nullTypes[mType].pkg,
		unsigned: goUtype{
			name: gType,
		},
	}
}

var types = map[string]goType{
	"time": {
		name: "string",
	},
	"date": {
		name: "dtype.Date",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"datetime": {
		name: "dtype.Time",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"timestamp": {
		name: "dtype.Time",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"year": {
		name: "string",
	},
	"decimal": {
		name: "decimal.Decimal",
		pkg:  "github.com/shopspring/decimal",
	},
	"numeric": {
		name: "decimal.Decimal",
		pkg:  "github.com/shopspring/decimal",
	},
	"bit": {
		name: "dtype.Bit",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"bits": {
		name: "dtype.Bits",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"char": {
		name: "string",
	},
	"varchar": {
		name: "string",
	},
	"tinytext": {
		name: "string",
	},
	"text": {
		name: "string",
	},
	"mediumtext": {
		name: "string",
	},
	"longtext": {
		name: "string",
	},
	"enum": {
		name: "string",
	},
	"set": {
		name: "dtype.Set",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"tinyint": {
		name:     "int8",
		unsigned: goUtype{name: "uint8"},
	},
	"smallint": {
		name:     "int16",
		unsigned: goUtype{name: "uint16"},
	},
	"mediumint": {
		name:     "int32",
		unsigned: goUtype{name: "uint32"},
	},
	"int": {
		name:     "int32",
		unsigned: goUtype{name: "uint32"},
	},
	"integer": {
		name:     "int32",
		unsigned: goUtype{name: "uint32"},
	},
	"bigint": {
		name:     "int64",
		unsigned: goUtype{name: "uint64"},
	},
	"double": {
		name: "float64",
	},
	"float": {
		name: "float64",
	},
	"real": {
		name: "float64",
	},
}

var nullTypes = map[string]goType{
	"date": {
		name: "dtype.NullDate",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"datetime": {
		name: "dtype.NullTime",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"timestamp": {
		name: "dtype.Time",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"time": {
		name: "sql.NullString",
		pkg:  "database/sql",
	},
	"year": {
		name: "sql.NullString",
		pkg:  "database/sql",
	},
	"bit": {
		name: "dtype.NullBit",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"bits": {
		name: "dtype.NullBits",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"char": {
		name: "sql.NullString",
		pkg:  "database/sql",
	},
	"varchar": {
		name: "sql.NullString",
		pkg:  "database/sql",
	},
	"tinytext": {
		name: "sql.NullString",
		pkg:  "database/sql",
	},
	"text": {
		name: "sql.NullString",
		pkg:  "database/sql",
	},
	"mediumtext": {
		name: "sql.NullString",
		pkg:  "database/sql",
	},
	"longtext": {
		name: "sql.NullString",
		pkg:  "database/sql",
	},
	"enum": {
		name: "sql.NullString",
		pkg:  "database/sql",
	},
	"set": {
		name: "dtype.NullSet",
		pkg:  "github.com/whf-sky/efficient/widget/dtype",
	},
	"double": {
		name:     "sql.NullFloat64",
		pkg:      "database/sql",
	},
	"float": {
		name: "sql.NullFloat64",
		pkg:  "database/sql",
	},
	"real": {
		name: "sql.NullFloat64",
		pkg:  "database/sql",
	},
	"decimal": {
		name: "decimal.Decimal",
		pkg:  "github.com/shopspring/decimal",
	},
	"numeric": {
		name: "decimal.Decimal",
		pkg:  "github.com/shopspring/decimal",
	},
	"tinyint": {
		name:     "dtype.NullInt8",
		pkg:      "github.com/whf-sky/efficient/widget/dtype",
		unsigned: goUtype{
			name:     "dtype.NullUint8",
			pkg:      "github.com/whf-sky/efficient/widget/dtype",
		},
	},
	"smallint": {
		name:     "sql.NullInt16",
		pkg:      "database/sql",
		unsigned: goUtype{
			name:     "dtype.NullUint16",
			pkg:      "github.com/whf-sky/efficient/widget/dtype",
		},
	},
	"mediumint": {
		name:     "sql.NullInt32",
		pkg:      "database/sql",
		unsigned: goUtype{
			name:     "dtype.NullUint32",
			pkg:      "github.com/whf-sky/efficient/widget/dtype",
		},
	},
	"int": {
		name: "sql.NullInt32",
		pkg:  "database/sql",
		unsigned: goUtype{
			name:     "dtype.NullUint32",
			pkg:      "github.com/whf-sky/efficient/widget/dtype",
		},
	},
	"integer": {
		name:     "sql.NullInt32",
		pkg:      "database/sql",
		unsigned: goUtype{
			name:     "dtype.NullUint32",
			pkg:      "github.com/whf-sky/efficient/widget/dtype",
		},
	},
	"bigint": {
		name: "sql.NullInt64",
		pkg:  "database/sql",
		unsigned: goUtype{
			name:     "dtype.NullUint64",
			pkg:      "github.com/whf-sky/efficient/widget/dtype",
		},
	},
}
