package main

import "fmt"

type car struct {
	color string
	name string
}

func main() {
	// Example 1
	car1 := car{"red", "220d"}
	car2 := new(car)
	car2.color, car2.name = "white", "sonata"
	car3 := &car{}
	car4 := &car{"black", "520d"}

	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car3)
	fmt.Println(car4)

}
