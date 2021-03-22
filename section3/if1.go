package main

import "fmt"

func main() {
	// 제어문(조건문)
	// IF 문은 반드시 Boolean 형으로 검사해야 한다. -> 1, 0, (사용불가 : 자동 형 변환 불가)
	// 소괄호 사용하지 않는다.
	var a int = 20
	b := 20

	// 예제 1
	if a >= 15 {
		fmt.Println("15 이상이다")
	}

	// 예제 2
	if b >= 25 {
		fmt.Println("25 이상이다")
	}

	// 에러 발생 1
	/*if b >= 25
	{

	}*/

	// 에러 발생 2
	/*if b>= 25
		fmt.Println("25이상")*/

	// 에러 발생 3
	if c := true; c {
		fmt.Println("ture")
	}

	// 예제 3
	if c := 40; c >= 35 {
		fmt.Println("35 이상")
	}
}
