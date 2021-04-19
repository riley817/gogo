package main

import "fmt"

type Account struct {
	number string
	balance float64
	interest float64
}

func (a Account) CalculateD(bonus float64) {
	a.balance = (a.balance + (a.balance * a.interest)) + bonus
}

func (a *Account) CalculateF(bonus float64) {
	a.balance = (a.balance + (a.balance * a.interest)) + bonus
}

func main() {
	// 정리 : 구조체 인스턴스 값 변경 시 -> 포인터 전달, 보통의 경우 -> 값 전달


	// Example 1
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-902", 20000000, 0.025}

	fmt.Println("Example 1 : ", kim)
	fmt.Println("Example 1 : ", lee)
	fmt.Println()

	kim.CalculateF(1000000)
	lee.CalculateD(1000000)

	fmt.Println("Example 1 : ", int(kim.balance))
	fmt.Println("Example 1 : ", int(lee.balance))
}
