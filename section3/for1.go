package main

import "fmt"

func main() {
	// 반복문 - For
	// Go 에서 유일하게 제공되는 반복문
	// 다양한 사용법을 숙지

	// Example 1
	for i := 0; i < 5; i++ {
		fmt.Println("EX1 : ", i)
	}

	// Error Case 1
	/*for i := 0; i < 5; i++
	{

	}*/

	// Error Case 2
	/*for i := 0; i < 5; i++
		fmt.Println("EX1 : ", i)*/

	// Example 2 - Infinite Loop
	/*for {
		fmt.Println("EX2 : Hello Go Lang ")
		fmt.Println("EX2 : Infinite Loop! ")
	}*/

	// Example 3 - Range 용법
	loc := []string {"Seoul", "Busan", "Incheon"}
	for index, name := range loc { // 첫번째 인자 - 인덱스, 두번째 인자 - 값
		fmt.Println("EX 3 : ", index, name)
	}

	for _, name := range loc {
		fmt.Println("EX 3 : ", name)
	}
}
