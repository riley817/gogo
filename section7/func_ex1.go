package main

import "fmt"

func multiply(n ...int) int {
	tot := 1
	for _, value := range n {
		tot *= value
	}
	return tot
}

func sum(n ...int) int {
	tot := 0
	for _, value := range n {
		tot += value
	}
	return tot
}

func prtWord(msg ...string) {
	for _, value := range msg {
		fmt.Println("Print Word ->", value)
	}
}


func main() {

	// 함수 고급
	// 가변 인자 실습 : 매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음.

	// Example 1
	x := multiply(5, 6, 7, 8, 9, 10)
	y := sum(1,2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13)
	fmt.Println("x -> ", x)
	fmt.Println("y -> ", y)
	fmt.Println()

	// Example 2
	prtWord("A", "F", "test", "golang", "seoul")

	// Example 3
	a := []int{5, 6, 7, 8, 9, 10}

	m := multiply(a...)
	n := sum(a...)

	fmt.Println()
	fmt.Println("Slice Parameter ->", m)
	fmt.Println("Slice Parameter ->", n)

}
