package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// 문자열
	// 큰 따옴표 "", 백스쿼트 ``
	// GoLang : 문자의 char 타입이 존재 하지 않는다. -> rune(int32) 로 문자 코드 값으로 표현
	// 문자 : '' 로 작성 가능하다.
	// 자주 사용하는 escape : \\, \', \", \a(콘솔벨), \b(백스페이스), \f(쪽바꿈), \n(줄바꿈), \r(복귀), \t(탭) ... - 역슬래시 이후에 등장하는 문자를 그대로 출력할 때

	// Example 1
	var str1 string = "c:\\go_study\\src\\"	// c:\go_study\src\
	str2 := `c:\go_study\src\` 				// escape 를 사용하지 않아도

	fmt.Println("Example 1 : ", str1)
	fmt.Println("Example 1 : ", str2)

	// Example 2
	var str3 string = "Hello, world!"
	var str4 string = "안녕하세요!"
	var str5 string = "\ud55c\uae00"

	fmt.Println("############################")
	fmt.Println("Example 2 : ", str3)
	fmt.Println("Example 2 : ", str4)
	fmt.Println("Example 2 : ", str5)

	// Example 3
	// 길이(Byte 수),
	fmt.Println("############################")
	fmt.Println("Example 3 : ", len(str3))	// 13 byte
	fmt.Println("Example 3 : ", len(str4)) // 16 byte

	// Example 4
	// 길이(실제 길이)
	fmt.Println("############################")
	fmt.Println("Example 4 : ", utf8.RuneCountInString(str3))
	fmt.Println("Example 4 : ", utf8.RuneCountInString(str4))
	fmt.Println("Example 4 : ", len([]rune(str4)))	// utf8 패키지를 사용하지 않고도 문자열 길이를 구할 수 있다.
}
