package panic

import (
	"fmt"
	"testing"
)

func Test_panic(t *testing.T) {
	defer func() { fmt.Println("打印前") }()
	defer func() { fmt.Println("打印中") }()
	defer func() { fmt.Println("打印后") }()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	for i := 5; i >= 0; i-- {
		res := 10 / i
		fmt.Println(res)
	}
}
