package main

import "fmt"

func main() {
	// 슬라이스 복사
	// copy(복사 대상, 원본)
	// 반드시 make 로 공간을 할당 후 복사 해야 한다.
	// 복사 된 슬라이스 값 변경해도 원본에도 영향은 없음

	// Example 1
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice2 := make([]int, 5)
	slice3 := []int{}

	copy(slice2, slice1)	// slice2 가 용량이 5이므로 5만 복사
	copy(slice3, slice2)	// 복사안됨

	fmt.Println("Example 1 :", slice2)
	fmt.Println("Example 1 :", slice3)

	// Example 2
	a := []int{1, 2, 3, 4, 5}
	b := make([]int, 5)

	copy(b, a) // 가변길이로 배열처럼 사용할 수 있다.

	b[0] = 7
	b[4] = 10

	fmt.Println("Example 2 : ", a)
	fmt.Println("Example 2 : ", b)

	// Example 3
	c := [5]int{1, 2, 3, 4, 5}
	d := c[0:3]		// 주의! 부분적으로 슬라이스 추출은 참조 -> 원본 값이 변경된다.

	d[1] = 7
	fmt.Println()
	fmt.Println("Example 3 :", c)
	fmt.Println("Example 3 :", d)

	// Example 4
	e := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	f := e[0:5:7] // 용량 지정

	f[0] = 99

	fmt.Println(e)
	fmt.Println("Example 4 :", len(f), cap(f))
	fmt.Println("Example 4 :", f)


}