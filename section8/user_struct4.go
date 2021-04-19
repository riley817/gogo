package main

import "fmt"

type shoppingBasket struct { cnt, price int }

// 구매 함수
func (b shoppingBasket) purchase() int {
	fmt.Println(b.cnt)
	fmt.Println(b.price)
	fmt.Println()

	return b.cnt * b.price
}

// 원본 수정(참조 전달 형식)
func (b *shoppingBasket) rePurchaseP(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

// 원본 수정X (값 전달 형식)
func (b shoppingBasket) rePurchaseD(cnt, price int) {
	b.cnt += cnt
	b.price += price
}

func main() {
	// 리시버(구조체 메서드) 전달(값, 참조형식)
	// 함수는 기본적으로 값을 호출한다. -> 변수의 값이 복사 후 내부 전달(원본 수정 일어 나지 않음) -> 맵, 슬라이스 등은 참조 전달
	// 리시버(구조체)도 마찬가지로 포인터를 활용해서 메서드 내에서 원본 수정이 가능.

	// Example 1
	bs1 := shoppingBasket{3, 5000}
	fmt.Println("[Example 1] tot_price ", bs1.purchase())

	// 참조 전달 (원본 값 수정)
	//bs1.rePurchaseP(7, 5000)
	//fmt.Println("[Example 2] tot_price ", bs1.purchase())

	// 값 전달(원본 값 수정 x)
	bs1.rePurchaseD(3, 5000)
	fmt.Println("[Example 3] tot_price ", bs1.purchase())

}

