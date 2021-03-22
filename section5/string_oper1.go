package main

import "fmt"

// 데이터 타입 : 문자열 연산(1)
func main() {
	// 문자열 연산
	// 추출, 비교, 조합(결합)

	// Example 1 (추출)
	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("Example 1 : ", str1[0:2], str1[0])
	fmt.Println("Example 1 : ", str2[3:], str2[0])
	fmt.Println("Example 1 : ", str2[:4])
	fmt.Println("Example 1 : ", str1[1:3])
}
