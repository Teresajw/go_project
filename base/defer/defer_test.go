package _defer

import (
	"fmt"
	"testing"
)

func DeferClosures() {
	i := 0
	defer func() {
		println(i)
	}()
	i = 1
}

func DeferClosures1() {
	i := 0
	defer func(val int) {
		println(val)
	}(i)
	i = 1
}

// 指针
func DeferClosures2() {
	i := 0
	defer func(val *int) {
		println(*val)
	}(&i)
	i = 1
}

// defer 修改值
func DeferClosures3() int {
	a := 0
	defer func() {
		a = 1
	}()
	return a
}

func DeferClosures4() (a int) {
	a = 0
	defer func() {
		a = 1
	}()
	return
}

func DeferClosures5() {
	for i := 0; i < 10; i++ {
		//fmt.Printf("out, i的值是: %d , i的地址是：%p\n", i, &i)
		defer func() {
			fmt.Printf("in, i的值是: %d , i的地址是：%p\n", i, &i)
		}()
	}
}

func DeferClosures6() {
	for i := 0; i < 10; i++ {
		fmt.Printf("out, %p\n", &i)
		defer func(val int) {
			fmt.Printf("in, %p\n", &val)
			//println(val)
		}(i)
	}
}

func DeferClosures7() {
	for i := 0; i < 10; i++ {
		//fmt.Printf("out, i的地址：%p\n", &i)
		j := i
		fmt.Printf("out, j的地址：%p\n", &j)
		defer func() {
			fmt.Printf("in, j的地址：%p\n", &j)
			//println(j)
		}()
	}
}

func TestDeferClosures(t *testing.T) {
	//DeferClosures()
	//DeferClosures1()
	//DeferClosures2()
	//println(DeferClosures3())
	//println(DeferClosures4()) //只有返回值命名时才会写入

	DeferClosures5()
	//DeferClosures6()
	//DeferClosures7()

	/*go 1.21 之前，for 循环，共享变量*/
	/*go 1.22 之后，for 循环，拷贝副本*/
}
