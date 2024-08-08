package function

import "testing"

func Fun1(a, b int) int {
	// 函数式编程
	return a + b
}

func Test_Func(t *testing.T) {
	//a := Fun1
	//println(a(1, 2))
	Fun2()
	Fun3()
}

func Fun2() {
	fn := func() string {
		return "hello"
	}()
	println(fn)
}

func Fun3() {
	fn := func() string {
		return "111hello"
	}
	a := fn()
	println(fn)
	println(a)
}
