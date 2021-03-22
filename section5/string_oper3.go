package main

import (
	"fmt"
	"strings"
)

func main() {
	// Example 1 (결합 : 일반연산)
	// Java 와 비슷 - Java의 경우에도 문자열 일반연산보다 StrnigBuffer 쓰는거와 비
	str1 := "This document demonstrates the development of a simple Go package inside a module and introduces the go tool" +
		"the standard way to fetch, build, and install Go modules, packages, and commands.\n\n" +
		"Note: This document assumes that you are using Go 1.13 or later and the GO111MODULE environment variable is not set. If you are looking for the older," +
		"pre-modules version of this document, it is archived here."

	str2 := "This document demonstrates the development of a simple Go package"

	fmt.Println("Example ", str1 + str2)
	// Example 2 (결합 : Join) - *추천
	var strSet []string //슬라이스 선언
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("Example 2 :", strings.Join(strSet, "-----"))


}
