package main

import "fmt"

func main() {
	i := 7
	p := &i

	fmt.Println("Example 1 :", i, *p, &i, p)

	*p++	// 1증가
	fmt.Println("Example 1 :", i, *p, &i, p)

	*p = 7777 // 포인터 변수 역 참조 값 변경
	fmt.Println("Example 1 :", i, *p, &i, p)

	i = 77
	fmt.Println("Example 1 :", i, *p, &i, p)


}
