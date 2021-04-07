package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	name string "차랑명"
	color string "색상"
	company string "제조사"
}

func main() {
	// 필드 태그 사용

	// Example 1
	tag := reflect.TypeOf(Car{})

	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("Example 1 : ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}
}
