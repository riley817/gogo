package main

import (
	"fmt"
	"math"
)

// f의 i제곱 구하는 함수
func Power(f float64, n float64) (float64, error) {
	if f == 0 {
		return 0, fmt.Errorf("(%g)는/은 사용 불가능 합니다.", f)
	}
	return math.Pow(f, n), nil // 제곱, nil 반환
}

// 에러 처리 고급 - 1
func main() {
	// 에러 처리
	// Go error 패키지 New 메소드 사용 -> 사용자 에러 처리 생성

	// Example 1
	if f, err := Power(7, 3); err != nil {
		//fmt.Printf("Error Message : %s\n", err)
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("Example 1 : ", f)
	}

	// Example 2
	if f, err := Power(0, 3); err != nil {
		//fmt.Printf("Error Message : %s\n", err)
		fmt.Printf("Error Message : %s\n", err.Error())
	} else {
		fmt.Println("Example 2 : ", f)
	}
}
