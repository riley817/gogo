package main

import (
	"fmt"
	"section4/lib"
)

var num int32

// 변수 초기화
func init()  {
	num = 30
	fmt.Println("Main Init Start! => ", num)
}

func main()  {
	fmt.Println("10 보다 큰 수 : ", lib.CheckNum(num))
}


