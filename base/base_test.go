package main

import (
	"fmt"
	"math"
	"runtime"
	"strconv"
	"testing"
	"unsafe"
)

func Test_base(t *testing.T) {
	t.Log(runtime.GOARCH)
	t.Log(runtime.GOOS)
	t.Log(runtime.Version())
	t.Log(runtime.NumCPU())
	t.Log(runtime.NumGoroutine())
	t.Log(runtime.MemProfileRate)
	t.Log(strconv.IntSize)
}

func Test_base_1(t *testing.T) {
	var a uint = 7
	t.Logf("%b", a)

	t.Logf("%d\t%d", math.MaxInt8, math.MinInt8)
}

func Test_base_2(t *testing.T) {
	var a complex64 = 1 + 2i
	t.Logf("%T", a)
}

func Test_base_3(t *testing.T) {
	a := 1
	var p unsafe.Pointer = unsafe.Pointer(&a)
	var q = uintptr(p)
	t.Logf("%p", p)
	t.Logf("%x", q)
}

func Test_base_4(t *testing.T) {
	// 使用int64来确保可以存储2^63
	value := int64(1) << 63
	fmt.Println(value)
}
