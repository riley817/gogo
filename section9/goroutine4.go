package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// GoRoutine
	// 클로저 사용 예제

	runtime.GOMAXPROCS(1)
	s := "Go Routine Closure : "

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) // 반복문 클로저는 일반적으로 즉시 실행 BUT 고루틴 클로저는 가장 나중에 실행(반복문 종료 후 실행)
	}
	time.Sleep(5 * time.Second)
}
