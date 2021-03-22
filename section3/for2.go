package main

import "fmt"

func main() {
	// Example 1
	sum1 := 0
	for i := 0; i <= 100; i++ {
		sum1 += 1 // sum := sum + 1
	}

	fmt.Println("EX1 : ", sum1)

	// Example 2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		// j := i++ (X) GO 에서는 후치연산은 반환이 안된다.
	}
	fmt.Println("EX2 : ", sum2)

	// Example 3
	sum3, i := 0, 0
	for { // while 형태와 비
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("EX2 : ", sum3)

	// Example 4
	for i, j := 0, 0; i <= 10; i, j = i + 1, j + 10 {
		fmt.Println("EX4 : ", i, j)
	}

	// Error Case
	/*for i, j := 0, 0; i <= 10; i++, j += 10 {

	}*/
}
