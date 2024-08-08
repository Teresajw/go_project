package types

import "testing"

type User struct {
	Id   int64
	Name string
	Age  int
	Sex  int
	Addr string
}

func Test_user(t *testing.T) {
	var u1 = &User{}
	println(u1.Id)

	var uPtr *User
	println(uPtr.Id)
}
