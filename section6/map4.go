package main

import "fmt"

func main() {
	// Map
	// 맵 조회 할 경우에 주의 할 점
	// Example 1
	map1 := map[string]int {
		"apple" : 25,
		"banana" : 115,
		"orange" : 1115,
		"lemon" : 0,
	}

	value1 := map1["lemon"]
	value2, ok2 := map1["banana"]
	value3, ok := map1["kiwi"]

	fmt.Println(value1)
	fmt.Println(value2, ",", ok2)
	fmt.Println(value3, ",", ok) // 두 번째 리턴 값으로 키 존재 유무 확인

	// Example 2
	if value, ok := map1["kiwi"]; ok {
		fmt.Println("Example 2 : ", value)
	} else {
		fmt.Println("Example 2 : kiwi is not exists")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("kiwi is not exists")
	}



}
