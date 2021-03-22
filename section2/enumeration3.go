package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		_
		C
	)

	fmt.Println(A, C)

	const (
		_ = iota + 0.75 * 2
		DEFAULT
		SILVER
		_
		PLATINUM
	)

	fmt.Println("D : ", DEFAULT)
	fmt.Println("S : ", SILVER)
	// fmt.Println("G : ", GOLD) 	// 사용하는 코드는 수정해야 한다.
	fmt.Println("P : ", PLATINUM)

}
