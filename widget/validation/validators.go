package validation

import (
	"errors"
	"fmt"
	"github.com/whf-sky/efficient"
)

func Lang(key string) string {
	if _, ok := efficient.EConfig.Lang["validation_"+key]; ok {
		return efficient.EConfig.Lang["validation_"+key]
	}
	return efficient.EConfig.Lang["validation_no_lang"]
}

func Error(key string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(Lang(key), a...))
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
