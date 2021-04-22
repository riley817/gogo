package main

import (
	"github.com/tealeg/xlsx"
)

// 사용자 패키지 설치 및 활용 예제
func main() {
	// 외부 저장소 패키지 설치
	// 두 가지 방법
	// 1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
	// 2. go get 패키지 주소 설치 -> 사용

	// 선언 후 go get 설치 예제 - 엑셀 파일 읽기
	xfile := "section12/sample.xlsx"

	xlFile, err := xlsx.OpenFile(xfile)
	if err != nil {
		panic("Excel Loads Error!")
	}

	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Cols {
			for _, cell := range row.Cells {
				text := cell.st
			}
		}

	}


}

