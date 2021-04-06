package main

import "fmt"

// 자료형 - 슬라이스 심화
func main() {

	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	fmt.Printf("[Example 1] s1 => %d\n", cap(s1))

	s1 = append(s1, 6, 7)
	s2 = append(s1, s2...) 		// 슬라이스 뒤에 슬라이스 삽입할 경우 ... 사용
	s3 = append(s2, s3[0:3]...) // 추출 후 병
	//fmt.Printf("[Example 1] s1 => %d", cap(s1))

	fmt.Println("[Example 1] s1 =>", s1)
	fmt.Println("[Example 1] s2 =>", s2)
	fmt.Println("[Example 1] s3 =>", s3)

	// Example 2
	s4 := make([]int, 0, 5)
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("Example 2 -> len : %d, cap: %d, value : %v\n", len(s4), cap(s4), s4) // 길이 및 용량 자동 증가
	}
}