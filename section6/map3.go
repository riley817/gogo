package main

import "fmt"

// map
func main() {
	map1 := map[string]string {
		"daum" : "https://www.daum.net",
		"naver" : "https://www.naver.com",
		"google"  : "https://www.google.com",
		"home1" : "http://127.0.0.1",
	}

	fmt.Println("Example 1 : ", map1)

	map1["home2"] = "http://localhost" // 추가
	fmt.Println("Example 1 : ", map1)

	map1["home2"] = "http://localhost"
	fmt.Println("Example 1 :", map1)

	// Example 2 - delete
	delete(map1, "home2")
	fmt.Println("Example 1 :", map1)

}
