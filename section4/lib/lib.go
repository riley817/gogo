package lib

import "fmt"

// 변수 초기화
func init()  {
	fmt.Println("lib Package > init Start!!")
}

func CheckNum(c int32) bool {
	return c > 10
}

