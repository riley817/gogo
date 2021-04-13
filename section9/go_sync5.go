package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	// 고루틴 동기화 객체
	// 동기화 상태(조건) 메소드 사용
	// wait, notify, notifyAll(Java)
	// Wait, signal, Broadcast (GO)

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	ch := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			ch <- 777
			fmt.Println("GoRoutine Wait -> ", n)
			condition.Wait() // 고 루틴 멈춤(대기)
			fmt.Println("Waiting End", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		// <- ch
		fmt.Println("Received : ", <-ch )
	}

	/*for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Wake GoRoutine(Signal)", i)
		condition.Signal() // 모든 고 루틴 생성 후 한개 씩 깨우기
		mutex.Unlock()
	}*/

	mutex.Lock()
	fmt.Println("Wake Goroutine(Broadcast)")
	condition.Broadcast()
	mutex.Unlock()

	time.Sleep(2 * time.Second)
}


