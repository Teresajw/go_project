package function

import (
	"fmt"
	"testing"
)

// defer 语句用于注册一个延迟调用的函数。允许你在返回的前一刻执行一段逻辑

func df() func() string {
	i := 0
	j := 0
	k := 0
	return func() string {
		defer func(x int) {
			println("defer1", i, j, *&k)
		}(j)
		i++
		j++
		k++
		defer func(x int) {
			println("defer2", i, j, *&k)
		}(j)
		i++
		j++
		k++
		defer func(x int) {
			println("defer3", i, j, *&k)
		}(j)
		return "hello world"
	}
}

func DeferReturn() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferReturnV1() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return
}

func DeferLooperV1() {
	for i := 0; i < 10; i++ {
		defer func() {
			println("defer", i)
		}()
	}
}

func DeferLooperV2() {
	for i := 0; i < 10; i++ {
		defer func(val int) {
			println("defer", val)
		}(i)
	}
}

func DeferLooperV3() {
	for i := 0; i < 10; i++ {
		j := i
		defer func() {
			println(j)
		}()
	}
}

func Test_defer(t *testing.T) {
	//t.Log(df()())
	//println(DeferReturn())
	//println(DeferReturnV1())
	//DeferLooperV1()
	//DeferLooperV2()
	DeferLooperV3()
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func Test_calc(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1

}
