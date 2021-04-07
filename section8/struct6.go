package main

import "fmt"

type Car struct {	// 첫글자 대문자로 선언
	name string "차랑명"
	color string "색상"
	company string "제조사"
	detail spec "상세"
}

type spec struct { // 첫글자 소문자로 선언
	length int "전장"
	height int "전고"
	width int "전축"
}

func main() {
	// 중첩 구조체
	car1 := Car{
		"520d",
		"silver",
		"bmw",
		spec{4000, 1000, 2000},
	}
	// 내부 spec 구조체 값 출력
	fmt.Println("Example 1 : ", car1.name)
	fmt.Println("Example 1 : ", car1.color)
	fmt.Println("Example 1 : ", car1.company)
	fmt.Printf("Example 1 : %#v\n", car1.detail)

	fmt.Println()
	fmt.Println("Example 1 : ", car1.detail.length)
	fmt.Println("Example 1 : ", car1.detail.height)
	fmt.Println("Example 1 : ", car1.detail.width)



}
