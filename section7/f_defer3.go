package main

import "fmt"

// FILO
func stack() {
	for i := 1; i <= 10; i++ {
		defer fmt.Println("Example 1 : ", i)
	}
}

func main() {
	// Example 1.
	stack()
}
