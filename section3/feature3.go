package main

import "fmt"

func main() {
	// 코드 서식을 지정해주는 유틸을 내장하고 있다.
	// 한 사람이 코딩한 것 같은 일관성을 유지
	// 코드 스타일을 유지
	// gofmt -h : 사용법
	// gofmt -w : 원본 파일에 반영

	// Example 1
	for i := 0; i <= 100; i++ {
		fmt.Println("ex 1 : ", i)
	}
}
