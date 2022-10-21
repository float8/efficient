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
		name: "database.Date",
		pkg:  "github.com/float8/efficient/database",
	},
	"datetime": {
		name: "database.Time",
		pkg:  "github.com/float8/efficient/database",
	},
	"timestamp": {
		name: "database.Time",
		pkg:  "github.com/float8/efficient/database",
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
		name: "database.Bit",
		pkg:  "github.com/float8/efficient/database",
	},
	"bits": {
		name: "database.Bits",
		pkg:  "github.com/float8/efficient/database",
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
		name: "database.Set",
		pkg:  "github.com/float8/efficient/database",
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
		name: "database.NullDate",
		pkg:  "github.com/float8/efficient/database",
	},
	"datetime": {
		name: "database.NullTime",
		pkg:  "github.com/float8/efficient/database",
	},
	"timestamp": {
		name: "database.Time",
		pkg:  "github.com/float8/efficient/database",
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
		name: "database.NullBit",
		pkg:  "github.com/float8/efficient/database",
	},
	"bits": {
		name: "database.NullBits",
		pkg:  "github.com/float8/efficient/database",
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
		name: "database.NullSet",
		pkg:  "github.com/float8/efficient/database",
	},
	"double": {
		name: "sql.NullFloat64",
		pkg:  "database/sql",
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
		name: "database.NullInt8",
		pkg:  "github.com/float8/efficient/database",
		unsigned: goUtype{
			name: "database.NullUint8",
			pkg:  "github.com/float8/efficient/database",
		},
	},
	"smallint": {
		name: "sql.NullInt16",
		pkg:  "database/sql",
		unsigned: goUtype{
			name: "database.NullUint16",
			pkg:  "github.com/float8/efficient/database",
		},
	},
	"mediumint": {
		name: "sql.NullInt32",
		pkg:  "database/sql",
		unsigned: goUtype{
			name: "database.NullUint32",
			pkg:  "github.com/float8/efficient/database",
		},
	},
	"int": {
		name: "sql.NullInt32",
		pkg:  "database/sql",
		unsigned: goUtype{
			name: "database.NullUint32",
			pkg:  "github.com/float8/efficient/database",
		},
	},
	"integer": {
		name: "sql.NullInt32",
		pkg:  "database/sql",
		unsigned: goUtype{
			name: "database.NullUint32",
			pkg:  "github.com/float8/efficient/database",
		},
	},
	"bigint": {
		name: "sql.NullInt64",
		pkg:  "database/sql",
		unsigned: goUtype{
			name: "database.NullUint64",
			pkg:  "github.com/float8/efficient/database",
		},
	},
}
