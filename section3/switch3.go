package main

import "fmt"

func main() {
	// Example 1
	a := 30 / 15
	switch a {
	case 2, 4, 6: // i가 2, 4, 6 인 경우
		fmt.Println("a -> ", a, "는 짝수")
	case 1, 3, 5:
		fmt.Println("a -> ", a, "는홀수")
	}

	// Example 2
	// fallthrough -
	switch e := "go"; e {
	case "java":
		fmt.Println("Java!")
		fallthrough
	case "go":
		fmt.Println("go!")
		fallthrough
	case "python":
		fmt.Println("python!")
	case "ruby":
		fmt.Println("ruby!")
		fallthrough
	case "c":
		fmt.Println("c!")
	}
}
