package validation

import "regexp"

func Empty(comment string, data, param interface{}) error {
	s, ok := data.(string)
	if !ok {
		return ErrorType(comment)
	}
	if s == "" {
		return Error("empty", comment)
	}
	return nil
}

func Email(comment string, data, param interface{}) error {
	s, ok := data.(string)
	if !ok {
		return ErrorType(comment)
	}
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`
	reg := regexp.MustCompile(pattern)
	if !reg.MatchString(s) {
		return Error("email", comment)
	}
	return nil
}

func Regexp(comment string, data, param interface{}) error {
	d, ok := data.(string)
	if !ok {
		return ErrorType(comment)
	}
	p, ok := param.(string)
	if !ok {
		return ErrorType(comment)
	}
	reg := regexp.MustCompile(p)
	if !reg.MatchString(d) {
		return Error("regexp", comment)
	}
	return nil
}
