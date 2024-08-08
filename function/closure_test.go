package function

import (
	"fmt"
	"testing"
)

// 闭包
func closure(name string) func() string {
	return func() string {
		return "hello world," + name
	}
}

func Test_closure(t *testing.T) {
	cl := closure("test")
	println(cl())
	/*type args struct {
		name string
	}
	tests := []struct {
		name       string
		args       args
		wantReturn string
	}{
		{
			name:       "name 有值",
			args:       args{name: "test"},
			wantReturn: "hello world,test",
		},
		{
			name:       "name 空",
			args:       args{name: ""},
			wantReturn: "hello world,",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closure(tt.args.name); got() != tt.wantReturn {
				t.Errorf("closure() = %v, want %v", got(), tt.wantReturn)
			}
		})
	}*/
}

func c() func() int {
	var i = 0
	fmt.Printf("outer: %p\n", &i)
	return func() int {
		fmt.Printf("inner before: %p\n", &i)
		i++
		fmt.Printf("inner after: %p\n", &i)
		return i
	}
}

func Test_c(t *testing.T) {
	fn := c()
	println(fn())
	println(fn())
	println(fn())

	fn1 := c()
	println(fn1())
	println(fn1())
	println(fn1())
}

func getNames(name string, alias ...string) {
	for _, v := range alias {
		fmt.Println(name, v)
	}
}

func getNames_new(name string, alias ...any) {
	for _, v := range alias {
		fmt.Println(name, v)
	}
}

func Test_getName(t *testing.T) {
	getNames("test")
	getNames("test", "test1")

	a := []string{"里斯", "张三"}
	getNames("test", a...)
	//注意any
	getNames_new("test", a)
}
