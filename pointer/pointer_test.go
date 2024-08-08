package pointer

import (
	"fmt"
	"testing"
)

func Test_pointer(t *testing.T) {
	str := "hello"
	pStr := &str
	nStr := *pStr
	fmt.Println(str)
	fmt.Println(pStr)
	fmt.Println(nStr)
	str = "world"
	fmt.Println(str)
	fmt.Println(pStr)
	fmt.Println(nStr)
}
