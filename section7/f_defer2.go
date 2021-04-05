package main

import "fmt"

func sayHello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi !")
	}()
}

func main() {
	// Example 1
	sayHello("GoLang")
}
