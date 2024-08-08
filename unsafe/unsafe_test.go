package unsafe

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"
)

func PrintFieldOffset(entity any) {
	typ := reflect.TypeOf(entity)
	for i := 0; i < typ.NumField(); i++ {
		fd := typ.Field(i)
		fmt.Printf("%s: %d \n", fd.Name, fd.Offset)
	}
	fmt.Println("------------------")
}

func TestPrintFieldOffset(t *testing.T) {
	fmt.Println(unsafe.Sizeof(UserV1{}))
	PrintFieldOffset(UserV1{})

	fmt.Println(unsafe.Sizeof(UserV2{}))
	PrintFieldOffset(UserV2{})

	fmt.Println(unsafe.Sizeof(UserV3{}))
	PrintFieldOffset(UserV3{})
}

type UserV1 struct {
	Name    string
	age     int32
	Alias   []byte
	Address string
}

type UserV2 struct {
	Name    string
	age     int32
	age1    int32
	Alias   []byte
	Address string
}

type UserV3 struct {
	Name    string
	Alias   []byte
	Address string
	age     int32
}

type UserV4 struct {
	Name    string
	age     int32
	Alias   []byte
	Address string
}
