package main

import "fmt"

func rangeSum(rg int, c chan int) {
	sum := 0
	for i := 1; i <= rg; i++ {
		sum += i
	}
	c <- sum
}

func main() {
	// 채널(Channel)

	c := make(chan int)

	go rangeSum(1000, c)
	go rangeSum(7000, c)
	go rangeSum(5000, c)

	// 순서대로 데이터 수신 : 채널에서 값 수신 완료 될 때까지 대기
	result1 := <- c
	result2 := <- c
	result3 := <- c

	fmt.Println("Example 1 : ", result1)
	fmt.Println("Example 2 : ", result2)
	fmt.Println("Example 3 : ", result3)

}
