package main

import "fmt"

// 수신 전용 채널을 리턴값으로 리턴
func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)

	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}

func main() {
	// Channel
	// 채널 또한 함수의 반환 값으로 사용 가능하다.

	// Example 1
	c := sum(100)
	fmt.Println("Example 1 ",  <-c )

}
