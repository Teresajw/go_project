package main

import (
	"testing"
	"unicode/utf8"
)

func Test_string(t *testing.T) {
	var s string
	s = "你好hello"
	//t.Log(len(s))
	//t.Log(string(rune(123)))
	t.Log(len(s))
	t.Log(utf8.RuneCountInString(s))
}

func Test_bytes(t *testing.T) {
	var a byte = 'a'
	t.Logf("a is %c", a)

	var str string = "hello 我是字符"
	var b []byte = []byte(str)
	t.Log(b)
}

func Test_bool(t *testing.T) {
	a, b := true, false
	t.Log(!(a && b))
	t.Log(!(a || b))
}
