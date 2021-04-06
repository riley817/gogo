package main

import "fmt"

type cnt int

func main() {
	// 기본 자료형 사용자 정의 타입

	// Example 1
	a := cnt(5)
	fmt.Println("Example 1 :", a)

	// Example 2
	var b cnt = 15	// 보통은 이렇게 작성하는 것이 좋다.
	fmt.Println("Example 1 :", b)

	// Example 3
	// testConvertT(b) 		// 예외 발생(중요)
	testConvertT(int(b))	// 사용자 정의 타입 <-> 기본 타입 : 매개 변수 전달 시에 변환해야 사용 가능(int(변수))

	testConvertD(b)
}

func testConvertT(i int) {
	fmt.Println("Example 3 : (Default Type) : ", i)
}

func testConvertD(i cnt) {
	fmt.Println("Example 4 : (Custom Type) : ", i)
}
