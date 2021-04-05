package main

import "fmt"

func main() {
	// 맵(Map)
	// 맵 조회 및 순회(Iterator)

	// Example 1
	map1 := map[string]string {
		"daum" : "https://www.daum.net",
		"naver" : "https://www.naver.com",
		"google"  : "https://www.google.com",
	}

	fmt.Println("Example 1 : ", map1["google"])
	fmt.Println("Example 1 : ", map1["daum"])
	fmt.Println()

	// Example 2
	for k, v := range map1 {
		fmt.Println("Example 2 : ", k, v)
	}
	fmt.Println()

	for _, v := range map1 {
		fmt.Println("Example 2 : ", v)
	}
}