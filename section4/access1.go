package main

import (
	"fmt"
	"section4/lib2"
)

func main()  {
	// 패키지 접근 제어
	// 변수, 상수, 함수, 메서드, 구조체 등 식별자
	// 첫글자 대문자 : 패키지 외부에서 접근 가능
	// 첫글자 소문자 : 패키지 외부에서 접근 불가. 해당 패키지 내에서만 접근 가능

	fmt.Println("100 보다 큰 수 ?", lib2.CheckNum1(101))
	fmt.Println("1000 보다 큰 수 ?", lib2.CheckNum2(999))

}