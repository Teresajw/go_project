package str

import (
	"fmt"
	"testing"
)

func Func1() {
	fmt.Println("func1 运行了")
}

func Func2(a int, b int) {
	res := fmt.Sprintf("func2 执行了, a+b=%d", a+b)
	fmt.Println(res)
}

func Func3(a, b int) (res int) {
	res = a + b
	return
}

func Func4(a, b int) (jia, jian int) {
	jia = a + b
	jian = a - b
	return
}

func Test_func(t *testing.T) {
	Func1()
	Func2(1, 2)
	r := Func3(1, 5)
	fmt.Println(r)
	_, j := Func4(1, 5)
	fmt.Printf("a-b=%d\n", j)
}
