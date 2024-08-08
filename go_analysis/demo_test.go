package go_analysis

import "testing"

type User struct {
	Name string
}

func ReturnPointer() *User {
	return &User{Name: "test"}
}

func TestReturnPointer(t *testing.T) {
	ReturnPointer()
}
