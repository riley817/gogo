package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	// defer 함수 - panic이 호출 되면 실행
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error! -> ", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Filename : ", f.Name())
	}

	defer f.Close()
}

func main() {
	// Go 패닉 & 리커버
	// Example 1
	fileOpen("undefined.txt")
	fmt.Println("End Main.")
}
