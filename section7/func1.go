package main

import (
	"fmt"
	"strconv"
)

// 함수 선언 위치는 어느 곳이 든 가능
func helloGoLang() { // 매개변수 X, 반환값 x
	fmt.Println("Example 1 : Hello, GoLang")
}

func say_one(m string) {
	fmt.Println("Example 2 : ", m)
}

func sum(x int, y int) int {
	return x + y
}

// 함수 기초
func main() {
	// 선언 : func 키워드로 선언
	// func 함수명(매개변수) (반환타입 or 반환 값 변수명) : 반환 값 존재
	// func 함수명() (반환타입 or 반환 값 변수명) : 매개변수 없음, 반환 값 존재
	// func 함수명(매개변수) : 매개변수 존재, 반환 값 없음
	// func 함수명() : 매개변수 없음, 반환값 없음
	// 타 언어와 달리 변환 값(return value) 다수 가능

	// Example 1
	helloGoLang()

	// Example 2
	say_one("Hello World!")

	// Example 3
	result := sum(5, 5)
	fmt.Println("Example 3 : ", result)
	fmt.Println("Example 3 : ", sum(5, 5))
	fmt.Println("Example 3 : ", strconv.Itoa(sum(5,5)))

}

