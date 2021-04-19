package main

import (
	"fmt"
	"sync"
)

// 고루틴 동기화 고급 - 2
func main() {
	// 고루틴 동기화 고급
	// 대기 그룹
	// 실행 흐름을 변경 하는데(고루틴이 종료 될 때까지 대기 가능)
	// WaiteGroup : Add - 고루틴 추가, Done(작업종료 알림), Wait - 고루틴 종료시까지 대기
	// Add 로 추가된 고루틴 개수와 Done 으로 종료되는 알림 횟수가 같아야 정확하게 동작한다. (Add == Done)

	wg := new(sync.WaitGroup)

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(n int) {
			fmt.Println("WaitGroup => ", n)
			wg.Done()
		}(i)
	}

	// Add 와 Done 횟수가 같아야한다.
	wg.Wait() // 대기 그룹이 끝날떄까지 기다린다.
	fmt.Println("WaitGroup End!")
}
