package generics

import "testing"

// ~int 支持衍生类型
func Sum[T ~int | int64 | float64](val ...T) T {
	var result T
	for _, v := range val {
		result += v
	}
	return result
}

type MyInt int

func SumV1[T Number](val ...T) T {
	var result T
	for _, v := range val {
		result += v
	}
	return result
}

// 实数
type RealNumber interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 |
		~int | ~int8 | ~int32 | ~int64 |
		~float32 | ~float64
}

// 复数
type Number interface {
	RealNumber | ~complex64 | ~complex128
}

func TestSum(t *testing.T) {
	println(Sum(1, 2.0))
	println(Sum(1, 2, 3))
	println(Sum(1, 2, 3))
}
