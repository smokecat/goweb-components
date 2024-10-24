package xvalidator

import (
	"github.com/go-playground/validator/v10"
)

var (
	defaultValidator *validator.Validate
)

func SetDefault(v *validator.Validate) {
	if v != nil {
		defaultValidator = v
		return
	}

	defaultValidator = validator.New(validator.WithRequiredStructEnabled())
}

func Default() *validator.Validate {
	return defaultValidator
}

func RegisterValidation(v *validator.Validate, tag string, f validator.Func, msgFunc customTagMsgFunc) error {
	if v == nil {
		v = defaultValidator
	}
	if msgFunc != nil {
		customTagMsgMap[tag] = msgFunc
	}

	return v.RegisterValidation(tag, f)
}
