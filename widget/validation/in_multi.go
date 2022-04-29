package validation

import "strconv"

func InMulti(comment string, data, params interface{}) error {
	switch data.(type) {
	case int8:
		return InMultiInt8(comment, data.([]int8), params)
	case uint8:
		return InMultiUint8(comment, data.([]uint8), params)
	case int16:
		return InMultiInt16(comment, data.([]int16), params)
	case uint16:
		return InMultiUint16(comment, data.([]uint16), params)
	case int32:
		return InMultiInt32(comment, data.([]int32), params)
	case uint32:
		return InMultiUint32(comment, data.([]uint32), params)
	case int:
		return InMultiInt(comment, data.([]int), params)
	case uint:
		return InMultiUint(comment, data.([]uint), params)
	case int64:
		return InMultiInt64(comment, data.([]int64), params)
	case uint64:
		return InMultiUint64(comment, data.([]uint64), params)
	case float32:
		return InMultiFloat32(comment, data.([]float32), params)
	case float64:
		return InMultiFloat64(comment, data.([]float64), params)
	case string:
		return InMultiString(comment, data.([]string), params)
	}
	return nil
}

func InMultiInt8(comment string, data []int8, params interface{}) error {
	p, ok := params.([]int8)
	if !ok {
		return Error("param_type_error", comment)
	}
	dmap := map[int8]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[int8]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(int64(k), 10))
		}
	}
	return nil
}

func InMultiUint8(comment string, data []uint8, params interface{}) error {
	p, ok := params.([]uint8)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[uint8]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[uint8]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(int64(k), 10))
		}
	}
	return nil
}

func InMultiInt16(comment string, data []int16, params interface{}) error {
	p, ok := params.([]int16)
	if !ok {
		return Error("param_type_error", comment)
	}
	dmap := map[int16]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[int16]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(int64(k), 10))
		}
	}
	return nil
}

func InMultiUint16(comment string, data []uint16, params interface{}) error {
	p, ok := params.([]uint16)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[uint16]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[uint16]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(int64(k), 10))
		}
	}
	return nil
}


func InMultiInt32(comment string, data []int32, params interface{}) error {
	p, ok := params.([]int32)
	if !ok {
		return Error("param_type_error", comment)
	}
	dmap := map[int32]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[int32]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(int64(k), 10))
		}
	}
	return nil
}

func InMultiUint32(comment string, data []uint32, params interface{}) error {
	p, ok := params.([]uint32)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[uint32]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[uint32]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(int64(k), 10))
		}
	}
	return nil
}

func InMultiInt(comment string, data []int, params interface{}) error {
	p, ok := params.([]int)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[int]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[int]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(int64(k), 10))
		}
	}
	return nil
}

func InMultiUint(comment string, data []uint, params interface{}) error {
	p, ok := params.([]uint)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[uint]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[uint]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(int64(k), 10))
		}
	}
	return nil
}

func InMultiInt64(comment string, data []int64, params interface{}) error {
	p, ok := params.([]int64)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[int64]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[int64]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(k, 10))
		}
	}
	return nil
}

func InMultiUint64(comment string, data []uint64, params interface{}) error {
	p, ok := params.([]uint64)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[uint64]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[uint64]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatInt(int64(k), 10))
		}
	}
	return nil
}

func InMultiFloat32(comment string, data []float32, params interface{}) error {
	p, ok := params.([]float32)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[float32]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[float32]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatFloat(float64(k), 'f', 5, 32))
		}
	}
	return nil
}

func InMultiFloat64(comment string, data []float64, params interface{}) error {
	p, ok := params.([]float64)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[float64]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[float64]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, strconv.FormatFloat(k, 'f', 5, 32))
		}
	}
	return nil
}

func InMultiString(comment string, data []string, params interface{}) error {
	p, ok := params.([]string)
	if !ok {
		return Error("param_type_error", comment)
	}

	dmap := map[string]bool{}
	for _, v := range data {
		dmap[v] = true
	}

	pmap := map[string]bool{}
	for _, v := range p {
		pmap[v] = true
	}

	for k, _ := range dmap {
		if _, ok := pmap[k]; !ok{
			return InMultiError(comment, k)
		}
	}
	return nil
}

func InMultiError(comment, data string) error {
	return Error("in-multi", comment, data)
}
