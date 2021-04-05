package main

import "fmt"

func main() {
	// 함수 고급
	// 익명 함수 (Anonymous Functions)
	// 즉시 실행(선언과 동시에)

	// Example 1
	func() {
		fmt.Println("Example 1 : Anonymous Function !")
	}()

	// Example 2
	msg := "Hello GoLang!"

	func(m string) {
		fmt.Println("Example 2 : ", m)
	}(msg)

	// Example 3
	func(x, y int) {
		fmt.Println("Example 3 :", x + y)
	}(10, 20)

	// Example 4
	r := func(x, y int) int {
		return x * y
	}(10, 100)

	fmt.Println("Example 4 : ", r)
}