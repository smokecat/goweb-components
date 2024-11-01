package xutil

type ToFn[T, U any] func(t T) U

type SkipToFn[T, U any] func(t T, u U) bool

func SkipToNil[T any, U interface{}]() SkipToFn[*T, *U] {
	return func(t *T, u *U) bool {
		return u == nil
	}
}

func ToList[T, U any](tList []T, toFn ToFn[T, U], skipFn SkipToFn[T, U]) []U {
	uList := make([]U, 0, len(tList))
	for _, t := range tList {
		u := toFn(t)
		if skipFn != nil && skipFn(t, u) {
			continue
		}

		uList = append(uList, u)
	}
	return uList
}
