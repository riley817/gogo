package main

import "fmt"

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n - 1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("Example 2 : (", n, ")", " hi!")
	prtHello(n - 1)
}

func main() {
	// 함수 고급
	// 재귀 함수(Recursion)
	// 프로그램이 보기 쉽고, 코드 간결, 오류 수정 용이 : 장점
	// 코드 이해하기 어렵고, 기억공간을 많이 사용, 무한루프 가능성

	// Example 1
	x := fact(7)
	fmt.Println("Example 1 : ", x)

	// Example 2
	prtHello(10)
}
