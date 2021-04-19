package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

// 에러(예외) 처리 구조체
type PowError struct {
	time time.Time // 에러 발생 시간
	value float64  // 파라미터
	message string // 에러 메세지
	// value interface{}
}

func(e *PowError) Error() string {
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - %s", e.time, e.value, e.message)
}

func Power2(f, i float64) (float64, error) {
	if f == 0 {
		return 0, &PowError{time : time.Now(), value: f, message: "0은 사용할 수 없습니다."}
	}
	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다."}
	}
	if math.IsNaN(i) {
		return 0, &PowError{time: time.Now(), value: i, message: "숫자가 아닙니다."}
	}
	return math.Pow(f, i), nil
}

// Go 에러 처리 고급 - 3
func main() {

	// Example 1
	v, err := Power2(10, 3) // 정상
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Example 1 : ", v)

	// Example 2
	t, err := Power2(0, 3) // Error
	if err != nil {
		log.Fatal(err)
		fmt.Println(err.(*PowError).message)
	}
	fmt.Println("Example 1 : ", t)


}
