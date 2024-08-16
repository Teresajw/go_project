package slice

import "github.com/Teresajw/go_project/tools/internal/errs"

func Delete[T any](src []T, index int) ([]T, T, error) {
	length := len(src)
	if index < 0 || index > length-1 {
		var zero T
		return nil, zero, errs.NewErrIndexOutOfRange(length, index)
	}
	dst := make([]T, length-1)
	copy(dst, src[:index])
	copy(dst[index:], src[index+1:])
	return dst, src[index], nil
}

func FilterDelete[T any](src []T, m func(idx int, src T) bool) []T {
	emptyPos := 0
	for idx := range src {
		if m(idx, src[idx]) {
			continue
		}
		src[emptyPos] = src[idx]
		emptyPos++
	}
	return src[:emptyPos]
}
