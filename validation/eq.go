package validation

import "strconv"

func Eq(comment string, data, param interface{}) error {
	if err := TypeCompare(comment, data, param); err != nil {
		return err
	}
	switch data.(type) {
	case int8:
		return EqInt8(comment, data.(int8), param.(int8))
	case uint8:
		return EqUint8(comment, data.(uint8), param.(uint8))
	case int16:
		return EqInt16(comment, data.(int16), param.(int16))
	case uint16:
		return EqUint16(comment, data.(uint16), param.(uint16))
	case int:
		return EqInt(comment, data.(int), param.(int))
	case uint:
		return EqUint(comment, data.(uint), param.(uint))
	case int32:
		return EqInt32(comment, data.(int32), param.(int32))
	case uint32:
		return EqUint32(comment, data.(uint32), param.(uint32))
	case int64:
		return EqInt64(comment, data.(int64), param.(int64))
	case uint64:
		return EqUint64(comment, data.(uint64), param.(uint64))
	case float32:
		return EqFloat32(comment, data.(float32), param.(float32))
	case float64:
		return EqFloat64(comment, data.(float64), param.(float64))
	case string:
		return EqString(comment, data.(string), param.(string))
	case bool:
		return EqBool(comment, data.(bool), param.(bool))
	}
	return nil
}

func EqInt8(comment string, data, param int8) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func EqUint8(comment string, data, param uint8) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func EqInt16(comment string, data, param int16) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func EqUint16(comment string, data, param uint16) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func EqInt(comment string, data, param int) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func EqUint(comment string, data, param uint) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func EqInt32(comment string, data, param int32) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func EqUint32(comment string, data, param uint32) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func EqInt64(comment string, data, param int64) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(data, 10))
	}
	return nil
}

func EqUint64(comment string, data, param uint64) error {
	if data != param {
		return EqError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func EqFloat32(comment string, data, param float32) error {
	if data != param {
		return EqError(comment, strconv.FormatFloat(float64(param), 'f', 5, 32))
	}
	return nil
}

func EqFloat64(comment string, data, param float64) error {
	if data != param {
		return EqError(comment, strconv.FormatFloat(param, 'f', 5, 64))
	}
	return nil
}

func EqString(comment string, data, param string) error {
	if data != param {
		return EqError(comment, param)
	}
	return nil
}

func EqBool(comment string, data, param bool) error {
	if data != param {
		return EqError(comment, strconv.FormatBool(param))
	}
	return nil
}

func EqError(comment, param string) error {
	return Error("eq", comment, param)
}
