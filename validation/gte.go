package validation

import "strconv"

func Gte(comment string, data, param interface{}) error {
	if err := TypeCompare(comment, data, param); err != nil {
		return err
	}
	switch data.(type) {
	case int8:
		return GteInt8(comment, data.(int8), param.(int8))
	case uint8:
		return GteUint8(comment, data.(uint8), param.(uint8))
	case int16:
		return GteInt16(comment, data.(int16), param.(int16))
	case uint16:
		return GteUint16(comment, data.(uint16), param.(uint16))
	case int32:
		return GteInt32(comment, data.(int32), param.(int32))
	case uint32:
		return GteUint32(comment, data.(uint32), param.(uint32))
	case int:
		return GteInt(comment, data.(int), param.(int))
	case uint:
		return GteUint(comment, data.(uint), param.(uint))
	case int64:
		return GteInt64(comment, data.(int64), param.(int64))
	case uint64:
		return GteUint64(comment, data.(uint64), param.(uint64))
	case float32:
		return GteFloat32(comment, data.(float32), param.(float32))
	case float64:
		return GteFloat64(comment, data.(float64), param.(float64))
	case string:
		return GteString(comment, data.(string), param.(string))
	}
	return nil
}

func GteInt8(comment string, data, param int8) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GteUint8(comment string, data, param uint8) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}
func GteInt16(comment string, data, param int16) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GteUint16(comment string, data, param uint16) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}
func GteInt32(comment string, data, param int32) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GteUint32(comment string, data, param uint32) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GteInt(comment string, data, param int) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GteUint(comment string, data, param uint) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GteInt64(comment string, data, param int64) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(param, 10))
	}
	return nil
}

func GteUint64(comment string, data, param uint64) error {
	if data < param {
		return GteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GteFloat32(comment string, data, param float32) error {
	if data < param {
		return GteError(comment, strconv.FormatFloat(float64(param), 'f', 5, 32))
	}
	return nil
}

func GteFloat64(comment string, data, param float64) error {
	if data < param {
		return GteError(comment, strconv.FormatFloat(param, 'f', 5, 64))
	}
	return nil
}

func GteString(comment, data, param string) error {
	if data < param {
		return GteError(comment, param)
	}
	return nil
}

func GteError(comment, data string) error {
	return Error("gte", comment, data)
}
