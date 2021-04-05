package main

import "fmt"

func multiply(x, y int) (r1, r2 int) {
	r1 = x * 10
	r2 = y * 20
	return r1, r2
}

func multiply2(x, y int) (int, int) {
	return x * 10, y * 20
}

func main() {
	// 리턴 값 변수 사용
	a, b := multiply(10, 5)
	fmt.Println("Example 1 : ", a, b)

	// Example 2
	c, d := multiply2(10, 5)
	fmt.Println("Example 2 : ", c, d)
}