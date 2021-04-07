package main

import "fmt"

type Dog struct {
	name string
	weight int
}

// bite 메소드를 구현
func (d Dog) bite() {
	fmt.Println(d.name, " bites!")
}

// 동물의 행동 인터페이스 선언
type Behavior interface {
	bite()

}

func main() {
	// 인터페이스 구현 예제
	// Example 1
	dog1 := Dog{"poll", 10}

	var interface1 Behavior
	interface1 = dog1
	interface1.bite()

	// Example 2
	dog2 := Dog{"marry", 12}
	interface2 := Behavior(dog2)

	interface2.bite()

	// Example 3
	interface3 := []Behavior{dog1, dog2}

	// 인덱스 형태로 실행
	for idx, _ := range interface3 {
		interface3[idx].bite()
	}

	// 값 형태로 실행(인터페이스)
	for _, val := range interface3 {
		val.bite()
	}
}
