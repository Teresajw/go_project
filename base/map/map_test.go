package _map

import (
	"fmt"
	"testing"
)

func Test_Map(t *testing.T) {
	m := make(map[string]int, 10)
	m["a"] = 1
}

func Test_Map_1(t *testing.T) {
	var m map[string]int
	//m["a"] = 1
	t.Log(m)
}

func Test_Map_2(t *testing.T) {
	m := make(map[int][2]int)
	a := [2]int{1, 2}
	m[1] = a
	fmt.Println(m[1][1])
	if v, ok := m[1]; ok {
		fmt.Println(v[1], ok)
	} else {
		t.Log("not exist")
	}
	/*t.Log(m)
	m[1] = [2]int{1, 2}
	t.Log(m)*/
}
