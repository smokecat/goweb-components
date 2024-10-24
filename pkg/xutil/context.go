package xutil

import (
	"context"
)

func CtxValue[T any](ctx context.Context, key interface{}) (T, bool) {
	v, ok := ctx.Value(key).(T)
	return v, ok
}

func CtxValueQty[T any](ctx context.Context, key interface{}) T {
	v, ok := ctx.Value(key).(T)
	if ok {
		return v
	}

	var zero T
	return zero
}
