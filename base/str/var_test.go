package str

import "testing"

var aa = 10

const bb = 20

func Test_var_1(t *testing.T) {
	aa := 20
	t.Log(aa)
	//bb = 30 常量不能修改
	t.Log(bb)
}
