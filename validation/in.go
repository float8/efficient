package validation

import "strconv"

func In(comment string, data, params interface{}) error {
	switch data.(type) {
	case int8:
		return InInt8(comment, data.(int8), params)
	case uint8:
		return InUint8(comment, data.(uint8), params)
	case int16:
		return InInt16(comment, data.(int16), params)
	case uint16:
		return InUint16(comment, data.(uint16), params)
	case int32:
		return InInt32(comment, data.(int32), params)
	case uint32:
		return InUint32(comment, data.(uint32), params)
	case int:
		return InInt(comment, data.(int), params)
	case uint:
		return InUint(comment, data.(uint), params)
	case int64:
		return InInt64(comment, data.(int64), params)
	case uint64:
		return InUint64(comment, data.(uint64), params)
	case float32:
		return InFloat32(comment, data.(float32), params)
	case float64:
		return InFloat64(comment, data.(float64), params)
	case string:
		return InString(comment, data.(string), params)
	}
	return nil
}

func InInt8(comment string, data int8, params interface{}) error {
	p, ok := params.([]int8)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(int64(data), 10))
}

func InUint8(comment string, data uint8, params interface{}) error {
	p, ok := params.([]uint8)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(int64(data), 10))
}

func InInt16(comment string, data int16, params interface{}) error {
	p, ok := params.([]int16)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(int64(data), 10))
}

func InUint16(comment string, data uint16, params interface{}) error {
	p, ok := params.([]uint16)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(int64(data), 10))
}

func InInt32(comment string, data int32, params interface{}) error {
	p, ok := params.([]int32)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(int64(data), 10))
}

func InUint32(comment string, data uint32, params interface{}) error {
	p, ok := params.([]uint32)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(int64(data), 10))
}

func InInt(comment string, data int, params interface{}) error {
	p, ok := params.([]int)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(int64(data), 10))
}

func InUint(comment string, data uint, params interface{}) error {
	p, ok := params.([]uint)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(int64(data), 10))
}

func InInt64(comment string, data int64, params interface{}) error {
	p, ok := params.([]int64)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(data, 10))
}

func InUint64(comment string, data uint64, params interface{}) error {
	p, ok := params.([]uint64)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatInt(int64(data), 10))
}

func InFloat32(comment string, data float32, params interface{}) error {
	p, ok := params.([]float32)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatFloat(float64(data), 'f', 5, 32))
}

func InFloat64(comment string, data float64, params interface{}) error {
	p, ok := params.([]float64)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, strconv.FormatFloat(data, 'f', 5, 64))
}

func InString(comment, data string, params interface{}) error {
	p, ok := params.([]string)
	if !ok {
		return ErrorType(comment)
	}
	for _, v := range p {
		if data == v {
			return nil
		}
	}
	return InError(comment, data)
}

func InError(comment, data string) error {
	return Error("in", comment, data)
}
