package main

import "fmt"

func rangeSum(rg int) int {
	sum := 0
	for i := 1; i <= rg; i++ {
		sum += i
	}
	return sum
}

func main() {
	fmt.Println("Result : ", rangeSum(1000))
	fmt.Println("Result : ", rangeSum(7000))
	fmt.Println("Result : ", rangeSum(5000))

}
