package main

import "fmt"

func rptc(n *int) {
	*n = 77
}

func vptc(n int) {
	n = 77
}

func main() {
	// 포인터의 값 전달
	// 함수, 메서드 호출 시에 매개변수 값을 복사 전달 -> 함수, 메서드 내에서는 원복 값 변경 불가능
	// 원본 값 변경 위해서 포인터로 전달
	// 특히 크기가 큰 배열인 경우 값 복사시에 시스템에 부담 -> 포인터 전달로 해결(슬라이스, 맵 참조 전달이기 때문에) -> 배열도 포인터 전달로 참조전달 할 수있따

	// Example
	var a int = 10
	var b int = 10

	fmt.Println("Example 1 : ", a)
	fmt.Println("Example 1 : ", b)
	fmt.Println()

	rptc(&a)
	vptc(b)



}