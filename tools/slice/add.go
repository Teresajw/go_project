package slice

import (
	"github.com/Teresajw/go_project/tools/internal/errs"
)

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}
	// 创建一个新的切片，长度为原切片长度加1
	dst := make([]T, length+1)

	// 复制原切片的前index部分到新切片
	copy(dst, src[:index])

	// 在新切片的index位置插入新元素
	dst[index] = element

	// 复制原切片从index之后的部分到新切片index+1往后的位置
	copy(dst[index+1:], src[index:])
	// 返回新切片
	return dst, nil
}
