package validation

import "strconv"

func Lt(comment string, data, param interface{}) error {
	if err := TypeCompare(comment, data, param); err != nil {
		return err
	}
	switch data.(type) {
	case int8:
		return LtInt8(comment, data.(int8), param.(int8))
	case uint8:
		return LtUint8(comment, data.(uint8), param.(uint8))
	case int16:
		return LtInt16(comment, data.(int16), param.(int16))
	case uint16:
		return LtUint16(comment, data.(uint16), param.(uint16))
	case int32:
		return LtInt32(comment, data.(int32), param.(int32))
	case uint32:
		return LtUint32(comment, data.(uint32), param.(uint32))
	case int:
		return LtInt(comment, data.(int), param.(int))
	case uint:
		return LtUint(comment, data.(uint), param.(uint))
	case int64:
		return LtInt64(comment, data.(int64), param.(int64))
	case uint64:
		return LtUint64(comment, data.(uint64), param.(uint64))
	case float32:
		return LtFloat32(comment, data.(float32), param.(float32))
	case float64:
		return LtFloat64(comment, data.(float64), param.(float64))
	case string:
		return LtString(comment, data.(string), param.(string))
	}
	return nil
}

func LtInt8(comment string, data, param int8) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LtUint8(comment string, data, param uint8) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LtInt16(comment string, data, param int16) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LtUint16(comment string, data, param uint16) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LtInt32(comment string, data, param int32) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LtUint32(comment string, data, param uint32) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LtInt(comment string, data, param int) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LtUint(comment string, data, param uint) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LtInt64(comment string, data, param int64) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(param, 10))
	}
	return nil
}

func LtUint64(comment string, data, param uint64) error {
	if data >= param {
		return LtError(comment, strconv.FormatInt(int64(param), 10))
	}
	return nil
}

func LtFloat32(comment string, data, param float32) error {
	if data >= param {
		return LtError(comment, strconv.FormatFloat(float64(param), 'f', 5, 32))
	}
	return nil
}

func LtFloat64(comment string, data, param float64) error {
	if data >= param {
		return LtError(comment, strconv.FormatFloat(param, 'f', 5, 64))
	}
	return nil
}

func LtString(comment, data, param string) error {
	if data >= param {
		return LtError(comment, param)
	}
	return nil
}

func LtError(comment, param string) error {
	return Error("lt", comment, param)
}
