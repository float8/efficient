package validation

type lenCompareFunc func(d, p int) (string, bool)

func lenCompare(comment string, data, param interface{}, clFunc lenCompareFunc) error {
	d, ok := data.(string)
	if !ok {
		return Error("data_type_length_error", comment)
	}
	p, ok := param.(int)
	if !ok {
		return Error("param_type_noint_error", comment)
	}
	if key, ok := clFunc(len(d), p); !ok {
		return Error(key, comment, p)
	}
	return nil
}

func LenNeq(comment string, data, param interface{}) error {
	return lenCompare(comment, data, param, func(d, p int) (string, bool) {
		return "len-neq", d != p
	})
}

func LenEq(comment string, data, param interface{}) error {
	return lenCompare(comment, data, param, func(d, p int) (string, bool) {
		return "len-eq", d == p
	})
}

func LenGte(comment string, data, param interface{}) error {
	return lenCompare(comment, data, param, func(d, p int) (string, bool) {
		return "len-gte", d >= p
	})
}

func LenLte(comment string, data, param interface{}) error {
	return lenCompare(comment, data, param, func(d, p int) (string, bool) {
		return "len-lte", d <= p
	})
}

func LenGt(comment string, data, param interface{}) error {
	return lenCompare(comment, data, param, func(d, p int) (string, bool) {
		return "len-gt", d > p
	})
}

func LenLt(comment string, data, param interface{}) error {
	return lenCompare(comment, data, param, func(d, p int) (string, bool) {
		return "len-lt", d < p
	})
}
