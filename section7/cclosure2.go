package main

import "fmt"

func main() {
	// Example 1
	cnt := increaseCnt()

	fmt.Println("Count ->", cnt())
	fmt.Println("Count ->", cnt())
	fmt.Println("Count ->", cnt())
	fmt.Println("Count ->", cnt())
	fmt.Println("Count ->", cnt())

	anotherCnt := increaseCnt()

	fmt.Println("함수를 새로운 변수에 할당 시 초기화")
	fmt.Println("Another Count ->", anotherCnt())
	fmt.Println("Another Count ->", anotherCnt())
	fmt.Println("Another Count ->", anotherCnt())

}

func increaseCnt() func() int {
	n := 0 	// 지역변수(캡쳐됨)
	return func() int {
		n += 1
		return n
	}
}
