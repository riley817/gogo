package main

import "fmt"

func main() {
	// Go : 모호하거나, 애매한 문법 금지
	// 후치(후위) 연산자 허용 i++, 전치(전위) 연산자 비허용 ++i -> 문법에러
	// 증감연산 반환 값이 없음 sum := i++
	// 포인터(Pointer)를 허용. 연산은 비허용
	// 주석 - //, /**/ - 사용법 숙지

	// Example 1
	sum, i := 0, 0

	for i <= 100 {
		// sum += i++ // 예외 발생
		sum += i
		i++
		// ++i (전위증감 비허용)
	}
	fmt.Println("Example 1 : ", sum)
}