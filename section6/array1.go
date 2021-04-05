package main

import (
	"fmt"
)

func main() {
	// 배열
	// 배열은 용량, 길이 항상 같다.
	// 배열 vs 슬라이스 차이점 중요
	// 배열 : 길이가 고정되어 있다. / 슬라이스 : 길이가 가변이다.
	// 배열 : 값 타입 / 슬라이스 : 참조 타입
	// 배열 : 값을 복사 전달한다. / 슬라이스 : 참조 값 전달
	// 배열 : 전체 비교연산자 사용 가능 / 슬라이스 : 전체 비교 연산자 사용 불가
	// 대부분 슬라이스를 사용한다.

	// cap() : 배열, 슬라이스 용량
	// len() : 배열, 슬라이스 길이

	// Example 1
	var arr1 [5]int
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	var arr3 = [5]int{1, 2, 3, 4, 5}
	arr4 := [5]int{1, 2, 3, 4, 5}
	arr5 := [5]int{1, 2, 3}
	arr6 := [...]int{1, 2, 3, 4, 5} // 배열 길이를 확신 할 수 없을 때. 잘 사용하지 않음.
	arr7 := [5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	arr1[2] = 5

	fmt.Printf("%-5T %d %v\n", arr1, len(arr1), arr1) // %-5T 원 자료형
	fmt.Printf("%-5T %d %v\n", arr2, len(arr2), arr2)
	fmt.Printf("%-5T %d %v\n", arr3, len(arr3), arr3)
	fmt.Printf("%-5T %d %v\n", arr4, len(arr4), arr4)
	fmt.Printf("%-5T %d %v\n", arr5, len(arr5), arr5)
	fmt.Printf("%-5T %d %v\n", arr6, len(arr6), arr6)
	fmt.Printf("%-5T %d %v\n", arr7, len(arr7), arr7)

	// Example 2
	fmt.Println("#####################################")
	arr8 := [5]int{1, 2, 3, 4, 5}
	arr9 := [5]int{
		1,
		2,
		3,
		4,
		5,
	}
	arr10 := [...]string{"kim", "lee", "park"}

	fmt.Printf("%-5T %d %v\n", arr8, len(arr8), arr8)
	fmt.Printf("%-5T %d %v\n", arr9, len(arr9), arr9)
	fmt.Printf("%-5T %d %v\n", arr10, len(arr10), arr10)



}