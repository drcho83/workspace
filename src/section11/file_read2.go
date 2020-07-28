package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os" // 파일 읽고 쓰기는 기본적으로 os에서 제공
)

func errCheck1(e error) {
	if e != nil {
		panic(e)
	}
}

func errCheck2(e error) {
	if e != nil {
		fmt.Println(e)
		return
	}
}

func main() {
	// 파일 읽기
	//csv 파일 읽기 예제
	//파일 생성
	file, err := os.Open("sample.csv")
	//리소스 해제
	defer file.Close()

	//csv 리더 생성
	//rr := csv.NewReader(file)
	rr := csv.NewReader(bufio.NewReader(file)) // 권장
	//내용 읽기
	row, err := rr.Read()   //1개의 raw를 읽고
	row1, err1 := rr.Read() //1개의 raw를 읽고
	errCheck1(err)
	errCheck1(err1)
	fmt.Println("CSV Read Example")
	fmt.Println(row)
	fmt.Println(row[0], row[1], row[1:5])
	fmt.Println(row1[0], row1[1], row1[1:5])
	fmt.Println("=====================================================================")

	//한번에 읽어오기
	rows, err := rr.ReadAll()
	errCheck1(err)

	fmt.Println("CSV Read ALL")
	//fmt.Println(rows)
	fmt.Println(rows[5][1], " : ", rows[2][1], " : ", rows[6][1])
	fmt.Println("=====================================================================")

	fileInfo, err := file.Stat()

	fmt.Println("파일 정보 출력(1)", fileInfo, "\n")
	fmt.Println("파일 이름(2)", fileInfo.Name(), "\n")
	fmt.Println("파일 크기(3)", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간(4)", fileInfo.ModTime(), "\n")
	fmt.Println("=====================================================================")

	//Row 단위의 CSV 파일 읽기
	for i, row := range rows {
		//fmt.Println(i, row)
		for j := range row {
			fmt.Printf("%s ", rows[i][j])
		}
		fmt.Println()
	}

}
