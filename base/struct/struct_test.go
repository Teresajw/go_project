package _struct

import (
	"fmt"
	"testing"
)

type Student struct {
	Name  string
	Age   int
	Sex   string
	Score int
}

func Test_struct_1(t *testing.T) {
	s := Student{
		Name:  "张三",
		Age:   18,
		Sex:   "男",
		Score: 90,
	}

	ss := &Student{
		Name:  "张三",
		Age:   18,
		Sex:   "男",
		Score: 90,
	}

	fmt.Println(s)
	fmt.Println(ss)

	fmt.Printf("s的指针：%p\n", &s)
	fmt.Printf("ss的指针：%p\n", &ss)
	fmt.Printf("s => v => %v\n", s)
	fmt.Printf("s => +v => %+v\n", s)
	fmt.Printf("s => #v => %#v\n", s)
	fmt.Printf("s => T => %T\n", s)
}
