package main

import "fmt"

func main() {
	// Example 1
	cnt := increaseCnt()

	fmt.Println("Example 1 : ", cnt())
	fmt.Println("Example 1 : ", cnt())
	fmt.Println("Example 1 : ", cnt())
	fmt.Println("Example 1 : ", cnt())
	fmt.Println("Example 1 : ", cnt())

}

func increaseCnt() func() int {
	n := 0 	// 지역변수(캡쳐됨)
	return func() int {
		n += 1
		return n
	}
}
