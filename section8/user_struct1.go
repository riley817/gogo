package main

import (
	"fmt"
)

type Car struct {
	name string
	color string
	price int64
	tax int64
	weight int64
}

// 일반 메서드
func Price(c Car) int64 {
	return c.price + c.tax
}

// 구조체 <-> 메소드 바인딩
func (c Car) Price() int64 {
	return c.price + c.tax
}

// 사용자 정의 타입
func main() {
	// Go -> 객체 지향 타입을 구조체로 정의한다. (클래스, 상속 개념이 없다)
	// 객체 지향 - class (속성 : 멤버변수, 기능(상태 : 메소드)) -> 코드의 재 사용성, 코드의 관리가 용이, 신뢰성이 높은 프로그램
	// 객체지향 언어일까 ?
	// Go 는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스등을 통한 다형성 지원, 구조체를 통한 클래스 형태의 코딩 가능
	// 객체지향의 기본 개념 -> Go 에서 포함하고 있다. -> 객체 지향 프로그래밍 언어
	// 상태, 메소드를 분리해서 정의 - 결합성 없음
	// 사용자 정의 타입 : 구조체, 인터페이스, 기본 타입 (int, float, string ...), 함수
	// 구조체와 메소드를 연결을 통해서 타 언어의 클래스 형식 처럼 사용 가능 -> 객체 지향이다.

	// Example1
	bmw := Car{name : "520d", price : 500000000, color : "white", tax : 5000000}
	benz := Car{name : "220d", price : 600000000, color : "white", tax : 6000000}

	fmt.Println("Example 1: ", bmw, &bmw)
	fmt.Println("Example 1: ", benz, &benz)
	// fmt.Printf("Example 1: %p\n", &bmw)
	// fmt.Printf("Example 1: %p\n", &benz)

	// 이렇게는 사용하지 않음.
	fmt.Println("Example 2 :", Price(bmw))
	fmt.Println("Example 2 :", Price(benz))

	fmt.Println("Example 3 : ", bmw.Price())
	fmt.Println("Example 3 : ", benz.Price())

	fmt.Print("Example 4 :", &bmw == &benz)

}
