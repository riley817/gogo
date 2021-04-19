package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}


func main() {
	// 파일 읽기
	// CSV 파일 읽기 예제
	// 패키지 저장소를 통해서 Excel 등 다양한 파일 형식 쓰기, 읽기 가능
	// 패키지 Github 주소 : https://github.com/tealeg/xlsx
	// bufio : 파일이 용량이 클 경우 버퍼 사용 권장

	// 파일 생성
	file, err := os.Open("section11/sample.csv")
	errCheck(err)

	// 리소스 해제
	defer file.Close()

	// CSV Reader 생성
	//rr := csv.NewReader(file)
	rr := csv.NewReader(bufio.NewReader(file)) // 권장

	// CSV 내용 읽기
	row, err := rr.Read() // 1개의 Row 단위로 읽기
	errCheck(err)
	row2, err2 := rr.Read() // 1개의 Row 단위로 읽기
	errCheck(err2)

	fmt.Println("CSV Read Example")
	fmt.Println(row[0], row[1], row[2])
	fmt.Println(row2[0], row2[1], row2[2])
	fmt.Println(row[0:5])

	fmt.Println("=================================")

	rows, err := rr.ReadAll() // 전체 Row Read
	errCheck(err)

	fmt.Println("CSV Read All Example")
	fmt.Println(rows)
	fmt.Println(rows[5][1])

	// Row 단위로 CSV 파일 읽기
	for i, row := range rows {
		// fmt.Println(i, row)
		for j := range row {
			fmt.Printf("%s    ", rows[i][j])
		}
		fmt.Println()
	}

}