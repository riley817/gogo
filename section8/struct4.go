package main

import "fmt"

func main() {
	// 구조체 익명 선언 및 활용
	// type 구조체명 타입
	car1 := struct {name, color string}{"520d", "red"}

	fmt.Println("[Example 1] : ", car1)
	fmt.Printf("[Example 1] : %#v\n", car1)

	// Example 2
	cars := []struct{name, color string}{{"520d", "red"}, {"530i", "white"}, {"528i", "blue"}}

	for _, c := range cars {
		fmt.Printf("(%s, %s) ----- (%#v)\n", c.name, c.color, c)
	}

}
