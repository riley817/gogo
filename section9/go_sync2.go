package main

import (
	"fmt"
	"runtime"
	"sync"
)

// 구조체 선언 - 여러 Go Routine 에서 접근하는 공유 데이터
type count struct {
	num int
	mutex sync.Mutex	// 구조체 내부에 있기때문에
}

// 공유 데이터가 변경이 되는 부분
func (c *count) increment() {
	// 공유 데이터 수정 전 뮤텍스로 보호
	c.mutex.Lock()
	c.num += 1
	// 공유 데이터 수정 후 보호 해제
	c.mutex.Unlock()
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	// Go Routine 동기화 예제
	// 실행 흐름 제어 및 변수 동기화 가능
	// 공유 데이터 보호가 가장 중요
	// 뮤텍스(Mutex) -> 여러 고 루틴에서 작업하는 공유 데이터 보호
	// sync.Mutex 선언 후 Lock, Unlock 사용

	// 동기화 사용한 예제
	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())
	// runtime.GOMAXPROCS(1)

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // 다른 CPU 에게 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<- done
	}

	c.result()
}


