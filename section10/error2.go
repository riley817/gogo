package main

import (
	"fmt"
	"log"
)

// 메소드 리턴 값 error 타입 중요
func notZero(n int) (string, error)  {
	if n != 0 {
		s := fmt.Sprint("Hello GoLang : ", n)
		return s, nil
	}
	return "", fmt.Errorf("%d를 입력했습니다. Error ", n)
}


// Go 에러처리 기초 - 2
func main() {
	// 에러 처리
	// Errorf를 이용한 에러 처리

	a, err := notZero(1)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("A => ", a)

	b, err := notZero(0)
	if err != nil {
		log.Fatal(err)
	}

	// Fata 이후 프로그램 종료 후 실행 중지
	fmt.Println("B => ", b)
	fmt.Println("End Error Handling!")
}
