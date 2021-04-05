package main

import "fmt"

func start(t string) string {
	fmt.Println("start : ", t)
	return t
}
func end(t string) {
	fmt.Println("end : ", t)
}

func a() {
	defer end(start("b")) // defer 문에 있는 end()만 적용됨. 중첩 함수 주의 ! -> 웬만하면 사용하지 말것.
	fmt.Println("in a")
}

func main() {
	// Example 1
	a()
}
