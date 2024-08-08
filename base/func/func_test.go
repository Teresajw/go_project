package _func

import (
	"fmt"
	"testing"
)

func Test_func(t *testing.T) {
	fn1 := func() string {
		return "Hello World"
	}
	fn2 := func() string {
		return "Hello World"
	}()
	fmt.Printf("fn1是一个方法, 方法地址: %p,此时没有调用,使用fn1()调用后,fn1()=\"%s\"\n", fn1, fn1())
	fmt.Printf("fn2是一个方法并且调用了，返回值给了fn2, fn2=\"%s\"", fn2)
}
