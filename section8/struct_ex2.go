package main

import "fmt"

type Account struct {
	number string
	balance float64
	interest float64
}

func CalculateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	// Example 1
	kim := Account{"245-901", 10000000, 0.015}
	lee := &Account{"245-902", 20000000, 0.025}

	fmt.Println("Example 1 : ", kim)
	fmt.Println("Example 1 : ", lee)
	fmt.Println()

	CalculateD(kim)
	CalculateP(lee)
	fmt.Println("Example 1 : ", int(kim.balance))
	fmt.Println("Example 1 : ", int(lee.balance))

}
