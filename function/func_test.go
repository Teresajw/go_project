package function

import (
	"testing"
)

func Dg(start, end int) int {
	// 递归终止条件
	if start > end {
		return 0
	}
	// 递归调用，每次递归将问题规模缩小
	return start + Dg(start+1, end)
}

func Test_func(t *testing.T) {
	println(Dg(1, 100))
}
