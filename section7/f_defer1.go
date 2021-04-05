package main

import "fmt"

func ex_f2() {
	fmt.Println("f2 : called!")
}

func ex_f1() {
	fmt.Println("f1 : start!")
	defer ex_f2() // 마지막에 호출 된다.
	fmt.Println("f1 : end!")
}


// 함수 defer
func main() {
	// Defer 함수 실행(지연)
	// Defer 를 호출한 함수가 종료되기 직전에 호출
	// 타 언어의 finally 문과 비슷하다.
	// 주로 리소스를 반환, 열린 파일 닫기, Mutex 잠금 해제

	// Example 1
	ex_f1()
}
