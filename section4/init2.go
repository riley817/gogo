package main

import (
	"fmt"
	_ "section4/lib"
)

func init()  {
	// 다른 패키지에서 import 할 때 실행된다.
	fmt.Println("Init1 Method Start! ")
}

func init()  {
	fmt.Println("Init2 Method Start! ")
}

func init()  {
	fmt.Println("Init3 Method Start! ")
}

func init()  {
	fmt.Println("Init4 Method Start! ")
}

func main()  {
	fmt.Println("Main Method Start!")
}