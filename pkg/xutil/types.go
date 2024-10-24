package xutil

import (
	"reflect"
)

func Pointer[T comparable](v T) *T {
	return &v
}

func IsNil(v any) bool {
	if v == nil {
		return true
	} else if rv := reflect.ValueOf(v); rv.Kind() == reflect.Ptr && rv.IsNil() {
		return true
	}

	return false
}
