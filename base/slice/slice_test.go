package slice

import (
	"fmt"
	"testing"
)

func Test_slice(t *testing.T) {
	var s []int
	t.Logf("s:%+v, s是不是nil: %+v", s, s == nil)
	//数组
	a1 := [3]int{1, 2, 3}
	t.Logf("a1: %v, len:%d, cap:%d", a1, len(a1), cap(a1))
	a2 := a1[0:2]
	t.Logf("a2: %v, len:%d, cap:%d", a2, len(a2), cap(a2))
	a3 := [4]int{}
	t.Logf("a3: %v, len:%d, cap:%d", a3, len(a3), cap(a3))
}

func Test_slice_jiequ(t *testing.T) {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1, s2 := s[0:3], s[:3]
	t.Logf("s1: %v, s2:%v", s1, s2)
	s3, s4, s5 := s[0:], s[:], s[0:len(s)]
	t.Logf("s3: %v, s4:%v, s5:%v", s3, s4, s5)
}

func Test_slice_1(t *testing.T) {
	//切片
	s := make([]int, 10)
	t.Logf("初始s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	s = append(s, 10)
	t.Logf("变化s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice_2(t *testing.T) {
	s := make([]int, 0, 10)
	t.Logf("初始s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	s = append(s, 10)
	t.Logf("变化s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice_3(t *testing.T) {
	s := make([]int, 10, 11)
	t.Logf("初始s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	s = append(s, 10)
	t.Logf("变化s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
}

func Test_slice_4(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func Test_slice_5(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:9] //s[8:9] 的截取操作限定了 s1 的右边界，但这只是长度意义上的，对于容量，s1 仍然和 s 保持强关联性.
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}

func Test_slice_6(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1[0] = -1 //s1 指向的底层数组，在 s 变化时也会跟着变化.
	t.Logf("s: %v, s1: %v", s, s1)
}

func Test_slice_7(t *testing.T) {
	s := make([]int, 10, 12)
	v := s[10] //s[10] 越界了，会引发 panic.
	t.Logf("v: %d", v)
}

func Test_slice_8(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	s1 = append(s1, []int{10, 11, 12}...)
	v := s[10]
	t.Logf("s: %v, s1: %v, v: %d", s, s1, v)
}

func Test_slice_9(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	changeSlice_9(s1)
	t.Logf("s: %v", s)
}
func changeSlice_9(s1 []int) {
	s1[0] = -1
}

func Test_slice_10(t *testing.T) {
	s := make([]int, 10, 12)
	s1 := s[8:]
	fmt.Printf("s1 传入函数前: %p\n", s1)
	changeSlice_10(s1)
	fmt.Printf("s1 传入函数后: %p\n", s1)
	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	t.Logf("s1: %v, len of s1: %d, cap of s1: %d", s1, len(s1), cap(s1))
}
func changeSlice_10(s1 []int) {
	fmt.Printf("s1 函数中,append 前: %p\n", s1)
	s1 = append(s1, []int{10, 11, 12}...)
	fmt.Printf("s1 函数中,append 后: %p\n", s1)
}

func Test_slice_11(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	//切片的容量是从切片的起始位置到底层数组的末尾的元素个数
	t.Logf("s[:3]: %v, len of s[:3]: %d, cap of s[:3]: %d", s[:3], len(s[:3]), cap(s[:3]))
	t.Logf("s[3:]: %v, len of s[3:]: %d, cap of s[3:]: %d", s[3:], len(s[3:]), cap(s[3:]))

	t.Logf("s: %v, len of s: %d, cap of s: %d", s, len(s), cap(s))
	t.Logf("s[:2]: %v, len of s[:2]: %d, cap of s[:2]: %d", s[:2], len(s[:2]), cap(s[:2]))

	s = append(s[:2], s[3:]...)
	t.Logf("s: %v, len: %d, cap: %d", s, len(s), cap(s))
	//v := s[4]
	//t.Logf("v: %v", v)

}

func Test_slice_bianli(t *testing.T) {
	s := []int{0, 1, 2, 3, 4}
	//for i 循环
	/*for i := 0; i < len(s); i++ {
		t.Logf("i: %d, s[i]: %v", i, s[i])
	}*/
	//for range 循环
	/*for i, v := range s {
		t.Logf("索引: %d, 值: %v", i, v)
	}*/
	//for {} 循环
	index := 0
	for {
		if index > len(s)-1 {
			break
		}
		t.Logf("index: %d, s[index]: %v", index, s[index])
		index++
	}

}

// 切片的Add方法
func SAdd1(s *[]int, n int) []int {
	return append(*s, n)
}

func SAdd2(s []int, n int) []int {
	return append(s, n)
}

func Test_SAdd(t *testing.T) {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4, 5)
	t.Logf("开始s: %v,len:%d,cap:%d", s, len(s), cap(s))
	s = SAdd1(&s, 5)
	//s = SAdd2(s, 5)
	t.Logf("调用后s: %v,len:%d,cap:%d", s, len(s), cap(s))
}

func Test_SAddd(t *testing.T) {
	s := make([]int, 0, 5)
	s1 := append(s, 1, 2, 3, 4, 5)
	t.Log(s)
	t.Log(s1)
}

func Test_ss(t *testing.T) {
	s := make([]int, 2, 10)
	t.Log(s)
	a := s[6:10] //显示指定了范围
	t.Log(a)
	b := s[6:] //不指定范围，默认到切片长度
	t.Log(b)
}

func Test_ss1(t *testing.T) {
	s := make([]int, 2, 10)
	println(s[3:])
}
