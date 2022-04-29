package validation

import "strconv"

func Gt(comment string, data, param interface{}) error {
	if err := TypeCompare(comment, data, param); err != nil {
		return err
	}
	switch data.(type) {
	case int8:
		return GtInt8(comment, data.(int8), param.(int8))
	case uint8:
		return GtUint8(comment, data.(uint8), param.(uint8))
	case int16:
		return GtInt16(comment, data.(int16), param.(int16))
	case uint16:
		return GtUint16(comment, data.(uint16), param.(uint16))
	case int32:
		return GtInt32(comment, data.(int32), param.(int32))
	case uint32:
		return GtUint32(comment, data.(uint32), param.(uint32))
	case int:
		return GtInt(comment, data.(int), param.(int))
	case uint:
		return GtUint(comment, data.(uint), param.(uint))
	case int64:
		return GtInt64(comment, data.(int64), param.(int64))
	case uint64:
		return GtUint64(comment, data.(uint64), param.(uint64))
	case float32:
		return GtFloat32(comment, data.(float32), param.(float32))
	case float64:
		return GtFloat64(comment, data.(float64), param.(float64))
	case string:
		return GtString(comment, data.(string), param.(string))
	}
	return nil
}

func GtInt8(comment string, data, param int8) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GtUint8(comment string, data, param uint8) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GtInt16(comment string, data, param int16) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GtUint16(comment string, data, param uint16) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GtInt32(comment string, data, param int32) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GtUint32(comment string, data, param uint32) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GtInt(comment string, data, param int) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GtUint(comment string, data, param uint) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GtInt64(comment string, data, param int64) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(param, 10))
	}
	return nil
}

func GtUint64(comment string, data, param uint64) error {
	if data <= param {
		return GtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func GtFloat32(comment string, data, param float32) error {
	if data <= param {
		return GtError(comment, strconv.FormatFloat(float64(param), 'f', 5, 32))
	}
	return nil
}

func GtFloat64(comment string, data, param float64) error {
	if data <= param {
		return GtError(comment, strconv.FormatFloat(param, 'f', 5, 64))
	}
	return nil
}

func GtString(comment, data, param string) error {
	if data <= param {
		return GtError(comment, param)
	}
	return nil
}

func GtError(comment, data string) error {
	return Error("gt", comment, data)
}
