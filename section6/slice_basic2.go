package main

import "fmt"

func main() {
	// 슬라이스(슬라이스 참조 타입 증명)

	// Example 1 - 배열
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7


	fmt.Println("Array Example 1 : ", arr1)
	fmt.Println("Array Example 1 : ", arr2)

	// Example 2 - 슬라이스
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice2[0] = 7

	fmt.Println("Slice Example 2 : ", slice1)
	fmt.Println("Slice Example 2 : ", slice2)

	// 슬라이스 예와 상황
	slice3 := make([]int, 50, 100)
	fmt.Println("Example 3 : ", slice3[4])
	// fmt.Println("Example 3 : ", slice3[50]) // Error ! 길이 만큼 초기화 된다.
	fmt.Println()

	// Example 4
	for i, v := range slice1 {
		fmt.Println("Example 4 : ", i, v)
	}


}