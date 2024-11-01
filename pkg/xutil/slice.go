package xutil

func FindFirst[T any](arr []T, f func(T) bool) (T, bool) {
	var empty T
	for _, v := range arr {
		if f != nil && f(v) {
			return v, true
		}
	}
	return empty, false
}
