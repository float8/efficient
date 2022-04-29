package validation

import (
	"strconv"
)

func Neq(comment string, data, param interface{}) error {
	if err := TypeCompare(comment, data, param); err != nil {
		return err
	}
	switch data.(type) {
	case int8:
		return NeqInt8(comment, data.(int8), param.(int8))
	case uint8:
		return NeqUint8(comment, data.(uint8), param.(uint8))
	case int16:
		return NeqInt16(comment, data.(int16), param.(int16))
	case uint16:
		return NeqUint16(comment, data.(uint16), param.(uint16))
	case int32:
		return NeqInt32(comment, data.(int32), param.(int32))
	case uint32:
		return NeqUint32(comment, data.(uint32), param.(uint32))
	case int:
		return NeqInt(comment, data.(int), param.(int))
	case uint:
		return NeqUint(comment, data.(uint), param.(uint))
	case int64:
		return NeqInt64(comment, data.(int64), param.(int64))
	case uint64:
		return NeqUint64(comment, data.(uint64), param.(uint64))
	case float32:
		return NeqFloat32(comment, data.(float32), param.(float32))
	case float64:
		return NeqFloat64(comment, data.(float64), param.(float64))
	case string:
		return NeqString(comment, data.(string), param.(string))
	case bool:
		return NeqBool(comment, data.(bool), param.(bool))
	}
	return nil
}

func NeqInt8(comment string, data, param int8) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func NeqUint8(comment string, data, param uint8) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func NeqInt16(comment string, data, param int16) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func NeqUint16(comment string, data, param uint16) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}
func NeqInt32(comment string, data, param int32) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func NeqUint32(comment string, data, param uint32) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func NeqInt(comment string, data, param int) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func NeqUint(comment string, data, param uint) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func NeqInt64(comment string, data, param int64) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(param, 10))
	}
	return nil
}

func NeqUint64(comment string, data, param uint64) error {
	if data == param {
		return NeqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func NeqFloat32(comment string, data, param float32) error {
	if data == param {
		return NeqError(comment, strconv.FormatFloat(float64(param), 'f', 5, 32))
	}
	return nil
}

func NeqFloat64(comment string, data, param float64) error {
	if data == param {
		return NeqError(comment, strconv.FormatFloat(param, 'f', 5, 64))
	}
	return nil
}

func NeqString(comment, data, param string) error {
	if data == param {
		return NeqError(comment, param)
	}
	return nil
}

func NeqBool(comment string, data, param bool) error {
	if data == param {
		return NeqError(comment, strconv.FormatBool(param))
	}
	return nil
}

func NeqError(comment, data string) error {
	return Error("neq", comment, data)
}
