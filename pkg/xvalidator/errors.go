package xvalidator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"

	"github.com/smokecat/goweb-components/pkg/xcode"
	"github.com/smokecat/goweb-components/pkg/xerr"
)

type customTagMsgFunc func(fe validator.FieldError) string

var (
	customTagMsgMap = make(map[string]customTagMsgFunc)
)

func WrapErr(err error) error {
	ve, ok := err.(validator.ValidationErrors)
	if !ok {
		if _, ok := err.(xerr.XErr); ok {
			return err
		}
		return xerr.New(xcode.CodeInvalidParams, err.Error())
	}

	msgs := make([]string, 0, len(ve))
	for _, fe := range ([]validator.FieldError)(ve) {
		msgs = append(msgs, msgForTag(fe))
	}

	return xerr.New(xcode.CodeInvalidParams, strings.Join(msgs, "; "))
}

func msgForTag(fe validator.FieldError) string {
	if f, ok := customTagMsgMap[fe.Tag()]; ok {
		return f(fe)
	}

	switch fe.Tag() {
	case "required":
		return fmt.Sprintf("Field `%s` is required", fe.Field())
	case "min":
		return fmt.Sprintf("Field `%s` must be greater than or equal to %s", fe.Field(), fe.Param())
	case "max":
		return fmt.Sprintf("Field `%s` must be less than or equal to %s", fe.Field(), fe.Param())
	case "len":
		return fmt.Sprintf("Field `%s`'s length must be %s", fe.Field(), fe.Param())
	case "required_if":
		params := strings.Split(fe.Param(), " ")
		return fmt.Sprintf("Field `%s` is required if the field `%s` is equal to `%s`", fe.Field(), params[0], params[1])
	default:
		return fe.Error()
	}
}
