package main

import "fmt"

func main() {
	// 배열 복사
	// 값 복사 확인 중요

	// Example 1
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	arr2 := arr1

	fmt.Println("Example 1 : ", arr1, &arr1)
	fmt.Println("Example 1 : ", arr2, &arr2)

	fmt.Printf("Example 1 : %p %v\n", &arr1, arr1)
	fmt.Printf("Example 1 : %p %v\n", &arr2, arr2)

}