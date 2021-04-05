package main

import "fmt"

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) {
	fmt.Println("Example 1 : ", a + b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i *= 10 // *i = *i * 10
}

func main() {
	// 함수(콜백) 전달, 참조 전달(call by reference), 값 전달(call by value)
	// Example 1
	sum(100, add) // 매개변수로 함수를 전달

	// Example 2
	// call by value
	a := 100
	multi_value(a)
	fmt.Println("Example 2 : ", a)

	// Example 3
	//reference by value
	b := 100
	multi_reference(&b)
	fmt.Println("Example 3 :", b)
}