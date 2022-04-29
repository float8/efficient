package validation

func TypeCompare(comment string, t1, t2 interface{}) error {
	if DataType(t1) != DataType(t2) {
		return Error("param_type_error", comment)
	}
	return nil
}

func DataType(param interface{}) string {
	switch param.(type) {
	case int8:
		return "int8"
	case uint8:
		return "uint8"
	case int:
		return "int"
	case uint:
		return "uint"
	case int64:
		return "int64"
	case uint64:
		return "uint64"
	case string:
		return "string"
	case float32:
		return "float32"
	case float64:
		return "float64"
	case bool:
		return "bool"
	default:
		return ""
	}
}
