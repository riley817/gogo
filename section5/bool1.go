package main

import "fmt"

func main()  {
	// bool (boolean) : 참, 거짓
	// 조건부 논리 연산자랑 주로 사용 : !, ||(or), &&(and)
	// 암묵적 형 변환 x : 0, Nil -> false 변환 없음

	// Example 1
	var b1 bool = true
	var b2 bool = false
	b3 := true

	// b4 := 1

	fmt.Println("############ Example 1 ############")
	fmt.Println("b1 : ", b1)
	fmt.Println("b2 : ", b2)
	fmt.Println("b3 : ", b3)

	fmt.Println("b3 == b3 :", b3 == b3)
	fmt.Println("b1 && b3 :", b1 && b3)
	fmt.Println("b1 || b2 :", b1 || b2)
	fmt.Println("!b1 :", !b1)

	// 암묵적인 형 변환은 일어나지 않는다.
	/*if b4 {
		fmt.Println("Example ", b4)
	}*/

}