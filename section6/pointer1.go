package main

import "fmt"

// 자료형 포인터 1
func main() {
	// 포인터
	// Go : 포인터를 지원(C)
	// 변수 지역성, 연속된 메모리 참조 ..., 힙, 스택
	// 파이썬(인터프리터), 자바(JRE), 컴파일러, 인터프리터
	// 포인터지원 x (파이썬, C#, JAVA 등)
	// 주소의 값은 직접 변경 불가능 - 잘못된 코딩으로 인한 버그 방지
	// *(애스터리스크)로 사용한다.
	// nil 로 초기화 (nil != 0)

	// Example 1
	var a *int 				// 방법1
	var b *int  = new(int) 	// 방법2

	fmt.Println(a)
	fmt.Println(b)

	i := 7

	fmt.Println("Example 1 : ", i, &i)

	a = &i // 주소값을 전달
	b = &i // 주소값을 전달

	fmt.Println("Example 1 : ", a, &i)
	fmt.Println("Example 1 : ", &a)	// i의 주소값이 저장되어있는 a의 주소값
	fmt.Println("Example 1 : ", *a)	// 역참조

	fmt.Println()

	fmt.Println("Example 1 : ", b, &i)
	fmt.Println("Example 1 : ", &b)	// i의 주소값이 저장되어있는 b의 주소값
	fmt.Println("Example 1 : ", *b)	// 역참조

	var c = &i
	d := &i
	*d = 10

	fmt.Println()
	fmt.Println("Example 1 : ", c, &i)
	fmt.Println("Example 1 : ", &c)	// i의 주소값이 저장되어있는 b의 주소값
	fmt.Println("Example 1 : ", *c)	// 역참조

	fmt.Println()
	fmt.Println("Example 1 : ", d, &i)
	fmt.Println("Example 1 : ", &d)	// i의 주소값이 저장되어있는 b의 주소값
	fmt.Println("Example 1 : ", *d)	// 역참조

}
