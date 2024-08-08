package main

import "testing"

func Test_const(t *testing.T) {
	const a = 1
	t.Log(a)

	const (
		name = iota
		name1
		name2
		name3 = 2
		name4 = iota
		name5
		name6
	)

	const (
		n = iota << 1
		n1
		n2
	)
}
