package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// 파일 IO (1)
func main() {
	// 파일 읽기, 쓰기 -> ioutil 패키지 활용
	// 더욱 편리하고 직관적으로 파일을 읽고 쓰기가 가능
	// 아래 메소드 확인 가능
	// WriteFile(), ReadFile(), ReadAll() 등 사용 가능
	// https://golang.org/pkg/io/ioutil/

	s := "Hello Golang!\n File Write Test!\n"

	// 파일 모드 (chmod, chown, chgrp) -> 퍼미션 제공
	// 읽기 : 4, 쓰기 : 2, 실행 : 1
	// 소유자, 그룹, 기타사용자 순서 (777)
	// https://golang.org/pkg/os/#FileMode

	// 파일 쓰기
	err := ioutil.WriteFile("test_write1.txt", []byte(s), os.FileMode(0644))
	errCheck(err)

	// 파일 읽기
	data, err := ioutil.ReadFile("section11/sample.txt")
	errCheck(err)

	fmt.Println("=============================================")
	fmt.Println(string(data))
	fmt.Println("=============================================")

}
