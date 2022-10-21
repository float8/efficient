package validation

import (
	"errors"
	"fmt"
	"github.com/float8/efficient/validation/lang"
)

func Lang(key string) string {
	if _, ok := lang.Lang["v_"+key]; ok {
		return lang.Lang["v_"+key]
	}
	return lang.Lang["v_no_lang"]
}

func Error(key string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(Lang(key), a...))
}

func ErrorType(s ...interface{}) error {
	return Error("error_type", s...)
}

func Required(comment string, data, param interface{}) error {
	if data != nil {
		return Error("required", comment)
	}
	return nil
}

func init() {
	RegisterValidation("Required", Required)

	RegisterValidation("neq", Neq)
	RegisterValidation("eq", Eq)
	RegisterValidation("gt", Gt)
	RegisterValidation("gte", Gte)
	RegisterValidation("lt", Lt)
	RegisterValidation("lte", Lte)
	RegisterValidation("in", In)
	RegisterValidation("in-multi", InMulti)

	RegisterValidation("len-eq", LenEq)
	RegisterValidation("len-neq", LenNeq)
	RegisterValidation("len-gte", LenGte)
	RegisterValidation("len-lte", LenLte)
	RegisterValidation("len-gt", LenGt)
	RegisterValidation("len-lt", LenLt)

	RegisterValidation("empty", Empty)
	RegisterValidation("email", Email)
	RegisterValidation("regexp", Regexp)
}
