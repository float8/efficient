package validation

import "strconv"

func Lte(comment string, data, param interface{}) error {
	if err := TypeCompare(comment, data, param); err != nil {
		return err
	}
	switch data.(type) {
	case int8:
		return LteInt8(comment, data.(int8), param.(int8))
	case uint8:
		return LteUint8(comment, data.(uint8), param.(uint8))
	case int16:
		return LteInt16(comment, data.(int16), param.(int16))
	case uint16:
		return LteUint16(comment, data.(uint16), param.(uint16))
	case int32:
		return LteInt32(comment, data.(int32), param.(int32))
	case uint32:
		return LteUint32(comment, data.(uint32), param.(uint32))
	case int:
		return LteInt(comment, data.(int), param.(int))
	case uint:
		return LteUint(comment, data.(uint), param.(uint))
	case int64:
		return LteInt64(comment, data.(int64), param.(int64))
	case uint64:
		return LteUint64(comment, data.(uint64), param.(uint64))
	case float32:
		return LteFloat32(comment, data.(float32), param.(float32))
	case float64:
		return LteFloat64(comment, data.(float64), param.(float64))
	case string:
		return LteString(comment, data.(string), param.(string))
	}
	return nil
}

func LteInt8(comment string, data, param int8) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LteUint8(comment string, data, param uint8) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LteInt16(comment string, data, param int16) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LteUint16(comment string, data, param uint16) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}


func LteInt32(comment string, data, param int32) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LteUint32(comment string, data, param uint32) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}


func LteInt(comment string, data, param int) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LteUint(comment string, data, param uint) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LteInt64(comment string, data, param int64) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(param, 10))
	}
	return nil
}

func LteUint64(comment string, data, param uint64) error {
	if data > param {
		return LteError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LteFloat32(comment string, data, param float32) error {
	if data > param {
		return LteError(comment, strconv.FormatFloat(float64(param), 'f', 5, 32))
	}
	return nil
}

func LteFloat64(comment string, data, param float64) error {
	if data > param {
		return LteError(comment, strconv.FormatFloat(param, 'f', 5, 64))
	}
	return nil
}

func LteString(comment, data, param string) error {
	if data > param {
		return LteError(comment, param)
	}
	return nil
}

func LteError(comment, param string) error {
	return Error("lte", comment, param)
}
