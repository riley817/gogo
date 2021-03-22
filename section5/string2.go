package main

import "fmt"

// 데이터 타입 : 문자열 (2)
func main() {
	// 문자열 표현
	// 문자열 : UTF-8 인코딩 (유니코드 문자 집합) -> 바이트 수 유의 - 한글 3Byte

	// Example 1
	var str1 string = "Golang"
	var str2 string = "World"
	var str3 string = "고프로그래밍"

	fmt.Println("Example 1 : ", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Println("Example 1 : ", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Println("Example 1 : ", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5])
	/*
		Example 1 :  71 111 108 97 110 103
		Example 1 :  87 111 114 108 100
		Example 1 :  234 179 160 237 148 132
	*/

	// Example 2
	fmt.Println("########################################")
	fmt.Printf("Example 2 : %c %c %c %c %c %c\n", str1[0], str1[1], str1[2], str1[3], str1[4], str1[5])
	fmt.Printf("Example 2 : %c %c %c %c %c\n", str2[0], str2[1], str2[2], str2[3], str2[4])
	fmt.Printf("Example 2 : %c %c %c %c %c %c\n", str3[0], str3[1], str3[2], str3[3], str3[4], str3[5]) // 한글깨짐

	conStr := []rune(str3)
	fmt.Printf("Example 2 : %c %c %c %c %c %c\n", conStr[0], conStr[1], conStr[2], conStr[3], conStr[4], conStr[5]) // 한글 정상 출력

	// Example 3
	fmt.Println("########################################")
	for i, char := range str1 {
		fmt.Printf("Example 3 : %c(%d)\t", char, i)
	}
	fmt.Println()

	fmt.Println("########################################")
	for i, char := range  str2 {
		fmt.Printf("Example 4 : %c(%d)\t", char, i)
	}

}