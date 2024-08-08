package str

import (
	"reflect"
	"strings"
	"testing"
	"unicode/utf8"
)

// 各种类型的零值
func Test_type(t *testing.T) {
	var i int
	t.Log(i)
	var f float64
	t.Log(f)
	var b bool
	t.Log(b)
	var s string
	t.Logf("结果:[%s]", s)
	var d = []int{1, 2, 3}
	t.Log(reflect.TypeOf(d))
	var e []int
	t.Log(reflect.TypeOf(e))

	var g = [1]int{1}
	t.Log(reflect.TypeOf(g))

	a := make([]int, 10)
	t.Log(reflect.TypeOf(a), len(a), cap(a))

	Gba()
	//gba()
}

// 字符串
func Test_string(t *testing.T) {
	//var s string
	//t.Log(s)
	a := "123"
	b := "张三"
	c := a + b
	t.Log(c)
	t.Log(len(a), len(b), utf8.RuneCountInString(b))

	var m = 8 // 默认是int类型
	var n uint = 8
	t.Log(m, n)
	//t.Log(m==n) 不同类型不能比较
}

func Test_string_1(t *testing.T) {
	a := `sakjdad
	dajdjada
	kdlakdlad
	kdladka`
	t.Log(a)

	b := "sasasas\nsakjda\ndadad"
	t.Log(b)
}

func Test_string_2(t *testing.T) {
	// string 长度
	a := "迪迦"
	t.Log(len(a))
	t.Log(len([]byte(a)))

	//使用utf8.RuneCountInString()
	t.Log(utf8.RuneCountInString(a))

	//rune类型，就是int32
	//type rune = int32

	var b rune = 'a'
	t.Log(b)

	isRune := strings.ContainsRune(a, '迪')

	t.Log(isRune)

}
