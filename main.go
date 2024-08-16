package main

import "fmt"

func safeDivide(a, b int) (int, error) {
	if b == 0 {
		// 触发panic
		panic("division by zero")
	}
	return a / b, nil
}

func StackOverFlow() {
	fmt.Println("stack over flow")
	StackOverFlow()
}

func main() {
	StackOverFlow()

	/*defer func() {
		// 使用recover捕获panic
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()

	// 调用可能会触发panic的函数
	result, err := safeDivide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("10 / 0 =", result)

	fmt.Println("11111111111111")*/
}
