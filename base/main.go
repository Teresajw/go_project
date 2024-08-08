package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 456
	var b int = 123
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a / b)
	fmt.Println(a * b)
	fmt.Println(a % b)
	fmt.Println(a ^ b)
	fmt.Println(a & b)
	//fmt.Println(a / 0)

	fmt.Println(math.MaxInt8)
	fmt.Println(math.Ceil(1.2))
}
