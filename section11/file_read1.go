package main

import (
	"fmt"
	"os"
)

// 에러 체크 방식 1
func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

// 에러 체크 방식 2
func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	// 파일 읽기
	// Open : 기존 파일 열기
	// Close : 리소스 닫기
	// Read, ReadAt : 파일 읽기
	// 각 운영체제 권한 주의(오류 메세지 확인)
	// 탐색 seek 중요

	// 파일 열기
	file, err := os.Open("section11/sample.txt")
	errCheck1(err)

	// 읽기 예제 1
	fileInfo, err := file.Stat() // 파일 사이즈 확인 위해 정보 획득
	errCheck2(err)
	fd1 := make([]byte, fileInfo.Size()) // 슬라이스에 읽은 내용 담는다.
	ct1, err := file.Read(fd1)

	fmt.Println("파일 정보 출력 1 => ", fileInfo, "\n")
	fmt.Println("파일 이름 1 => ", fileInfo.Name(), "\n")
	fmt.Println("파일 크기 1 => ", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간 1 =>", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업 완료 1 => (%d bytes) \n\n", ct1)
	fmt.Println(string(fd1))

	// 읽기 예제 2 (탐색 : Seek(offset))
	o1, err := file.Seek(20, 0) // 0 : 처음위치, 1: 현재위치 2: 마지막위치
	errCheck2(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)

	errCheck2(err)

	fmt.Printf("읽기 작업 완료 2 => (%d bytes) (%d ret)\n\n", ct2, o1)
	fmt.Println(string(fd2))

	// 읽기 예제 3
	o2, err := file.Seek(0, 0)
	errCheck2(err)

	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset
	errCheck1(err)

	fmt.Printf("읽기 작업 완료 2 => (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(string(fd3))

	defer file.Close()


}
