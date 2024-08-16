package slice

import "testing"

func TestDiffSet(t *testing.T) {
	var a = []int{1, 2, 3}
	var b = []int{4, 5, 6}
	diff := DiffSet(a, b)
	t.Log(diff)
}
