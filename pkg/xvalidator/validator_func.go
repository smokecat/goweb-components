package xvalidator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

func RegexpValidationFunc(reg *regexp.Regexp) validator.Func {
	return func(fl validator.FieldLevel) bool {
		return fl.Field().IsZero() || reg.MatchString(fl.Field().String())
	}
}

func RegexpValidationMsgFunc(pattern string) customTagMsgFunc {
	return func(fe validator.FieldError) string {
		return fmt.Sprintf("Field `%s` must match the pattern `%s`", fe.Field(), pattern)
	}
}
