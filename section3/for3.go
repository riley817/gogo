package main

import "fmt"

func main() {
	// Example 1
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("EX1 : ", i , j)
		}
	}

	// Example 2
	for i := 10; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("EX2 : ", i)
	}

Loop2:
	// 에러 발생(Loop 레이블 밑에 관련이 없는 소스코드가 있으면 )
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("EX 3 : ", i, j)
		}
	}
}
