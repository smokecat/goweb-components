package xutil

import (
	"time"

	"golang.org/x/exp/constraints"
)

func Seconds[T constraints.Integer](cnt T) time.Duration {
	return time.Duration(cnt) * time.Second
}

func SecondsNum[T constraints.Integer](d time.Duration) T {
	return T(d / time.Second)
}
