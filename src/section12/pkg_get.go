//사용자 패키지 설치 및 활용 예제

package main

import (
	"fmt"

	"github.com/tealeg/xlsx"
	//웹 URL 복사 후 CDM 창에서 D:\workspace\go_study\src\section12>go get github.com/tealeg/xlsx
)

func main() {
	//외부 저장소 패키지 설치
	//2가지 방법
	//1. import 선언 후 폴더 이동 후 go get 설치 -> 사용
	//2. go get 패키지 주소 설 -> 선언

	//선언 후 go get 설치 예제(엑셀 파일 읽기)
	xfile := "sample.xlsx"

	xlfile, err := xlsx.OpenFile(xfile)

	if err != nil {
		panic("excel loads error!!")
	}

	for _, sheet := range xlfile.Sheets {
		for _, row := range sheet.Rows {
			for _, Cell := range row.Cells {
				text := cell.String()
				fmt.Printf("%s\t", text)
			}
			fmt.Println()
		}
	}
}
