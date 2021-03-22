package main

import "fmt"

func main()  {
	// 논리연산자
	fmt.Println("############ Example 1 ############")
	fmt.Println("true && true : " , true && true)
	fmt.Println(true && false)
	fmt.Println(false && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(false || false)
	fmt.Println(!true)
	fmt.Println(!false)

	// Example 2 비교연산자
	num1 := 15
	num2 := 37
	fmt.Println(num1 < num2)
	fmt.Println(num1 > num2)
	fmt.Println(num1 >= num2)
	fmt.Println(num1 <= num2)
	fmt.Println(num1 == num2)
	fmt.Println(num1 != num2)
}