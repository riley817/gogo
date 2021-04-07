package main

import "fmt"

type Account struct {
	number string
	balance float64
	interest float64
}

// 구조체 인스턴스를 생성한 뒤 포인터를 리턴한다.
func NewAccount(number string, balance float64, interest float64) *Account { // 포인터 반환이 아닌 경우 값 복사
	return &Account{number, balance, interest}
}

func main() {
	// 구조체 생성자 패턴 예제

	// Example 1
	kim := Account{number : "245-901", balance: 10000000, interest: 0.015}

	var lee *Account = new(Account)
	lee.number = "245-902"	// getter, setter
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("Example 1 : ", kim)
	fmt.Println("Example 1 : ", lee)

	// Example2
	park := NewAccount("245-903", 1700000, 0.04)
	fmt.Println("Example 2 : ", park)



}
