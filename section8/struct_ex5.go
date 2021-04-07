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

func (e Executives) Calculate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

func main() {
	// 구조체 임베디드 메소드 오버라이딩 패턴

	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 1500000, 200000}

	// 임원
	ex1 := Executives{Employee{"lee", 5000000, 1000000}, 1000000}

	fmt.Println("Example 1 :", int(ep1.Calculate()))
	fmt.Println("Example 1 :", int(ep2.Calculate()))

	// Employee 구조체 통해서 메소드 호출
	fmt.Println("Example 2 : ", int(ex1.Calculate()))
	fmt.Println("Example 2 : ", int(ex1.Employee.Calculate() + ex1.specialBonus))

}
