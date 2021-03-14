package main

import "fmt"

func main() {
	// 상수
	// const 사용 초기화, 한 번 선언 후에는 값 변경 금지, 고정 된 값을 이용하여 관리
	const a string = "Test1"
	const b = "Test2"
	const c int32 = 10 * 10
	// const d = getHeight() // 함수 결과값을 할당하는 경우 예외 발생. 함수 사용할 수 없다.
	const e = 35.6
	const f = false
	/*
	   에러 발생이 되는 경우
	   const g string
	   g = "Test3" // 상수는 선언과 동시에 할당이 되어야 한다.
	*/

	fmt.Println("a : ", a, "b : ", b, "c : ", c, "e : ", e, "f : ", f)

}
