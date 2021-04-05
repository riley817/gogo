package main

import "fmt"

func main() {
	// 배열 순회
	// Example 1
	arr1 := [5]int{1, 10, 100, 1000, 10000}

	// len 길이 반복
	for i := 0; i < len(arr1); i++ {
		fmt.Println("Example 1 : ", arr1[i])
	}
	fmt.Println("#############################")

	arr2 := [5]int{7, 77, 777, 7777, 77777}
	// range
	for i, v := range arr2 {
		fmt.Println("Example 2 : ", i, v)
	}

	for _, v := range arr2 {
		fmt.Println("Example 2 : ", v)
	}
	for v := range arr2 {
		fmt.Println("Example 2 : ", v) // 인덱스가 출력 된다.
	}
}