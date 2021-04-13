package main

import "fmt"

func multiply(x, y int) (r int) {
	r = x * y
	return r
}

func sum(x, y int) (r int) {
	r = x + y
	return r
}

func main() {
	// 함수 고급
	// 함수를 변수에 할당

	// Example 1 - 슬라이스에 할당
	f := []func(int, int) int {multiply, sum}

	a := f[0](10, 10)
	b := f[1](10, 10)

	fmt.Println("슬라이스에 함수 할당 ->", a, f[0](10,10))
	fmt.Println("슬라이스에 함수 할당 ->", b, f[1](10,10))
	fmt.Println()

	// Example 2 - 변수에 할당
	var f1 func(int, int) int = multiply
	f2 := sum

	fmt.Println("변수에 함수 할당 ->", f1(10, 10))
	fmt.Println("변수에 함수 할당 ->", f2(10, 10))
	fmt.Println()

	// Example 3 - 맵에 할당
	m := map[string] func(int, int) int{
		"mul_func" : multiply,
		"sum_func" : sum,
	}

	fmt.Println("Map에 함수 할당", m["mul_func"](10, 10))
	fmt.Println("Map에 함수 할당", m["sum_func"](10, 10))
}
