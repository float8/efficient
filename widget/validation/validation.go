package validation

import (
	"strings"
)

type ValidationHandle func(comment string, data, param interface{}) error

var validations = map[string]ValidationHandle{}

func RegisterValidation(name string, validation ValidationHandle) {
	validations[strings.ToLower(name)] = validation
}

func NewValidation() *Validation {
	return &Validation{errors: map[string]error{}}
}

type V map[string]interface{}

type Validation struct {
	errors map[string]error
	error error
}

func (v *Validation) Validator(key string, data interface{}, validators V, comments ...string) *Validation {
	if len(validators) == 0 {
		return v
	}
	comment := key
	if len(comments) > 0 && comments[0] != "" {
		comment = comments[0]
	}

	for validator, param := range validators {
		validator = strings.ToLower(validator)
		if _, ok := validations[validator]; ok {
			if err := validations[validator](comment, data, param); err != nil {
				v.errors[key] = err
				v.error = err
			}
		}
	}
	return v
}

func (v *Validation) Errors() map[string]error {
	return v.errors
}

func (v *Validation) Error() error {
	return v.error
}
