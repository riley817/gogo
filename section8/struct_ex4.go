package main

import "fmt"

type Employee struct {
	name string
	salary float64
	bonus float64
}

type Executives struct {
	Employee // is a 관계
	specialBonus float64
}

func (e Employee) Calculate() float64 {
	return e.salary + e.bonus
}

func main() {
	// 구조체 임베디드 패턴
	// 다른 관점으로 메소드를 재 사용하는 장점 제공
	// 상속을 허용하지 않는 go 언어에서 메소드 상속 활용을 위한 패턴

	// Example 1
	// 직원
	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 1500000, 200000}

	// 임원
	ex1 := Executives{Employee{"lee", 5000000, 1000000}, 1000000}

	fmt.Println("Example 1 :", int(ep1.Calculate()))
	fmt.Println("Example 1 :", int(ep2.Calculate()))

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("Example 1 : ", int(ex1.Calculate() + ex1.specialBonus))

}
