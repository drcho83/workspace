package main

import (
	_ "bufio"
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
	//파일 읽기
	//Open : 기존 파일 열기
	//Read, ReadAt: 파일 읽기
	//각 운영체제 권한 주의 (오류 메시지 확인)
	//예외 처리 정말 중요!!
	//탐색 seek 중요
	//package github 주서: https://github.com/tealeg/xlsx

	//파일 읽기 예제
	//파일 열기
	file, err := os.Open("sample.txt")
	errCheck1(err)

	//파일 읽기 예제1
	fileInfo, err := file.Stat() // 파일 사이즈 확인 위해 정보 획득
	errCheck2(err)

	fd1 := make([]byte, fileInfo.Size()) //슬라이스에 읽은 내용 담는다.
	ct1, err := file.Read(fd1)           // 만들어지 슬라이스 크기에 읽으면서 저장됨

	fmt.Println("파일 정보 출력(1)", fileInfo, "\n")
	fmt.Println("파일 이름(2)", fileInfo.Name(), "\n")
	fmt.Println("파일 크기(3)", fileInfo.Size(), "\n")
	fmt.Println("파일 수정 시간(4)", fileInfo.ModTime(), "\n")
	fmt.Printf("읽기 작업(1) 완료 (%d bytes)\n\n", ct1)
	fmt.Println(fd1)
	fmt.Println(string(fd1))

	fmt.Println("=======================================")

	//seek(offset)
	ol, err := file.Seek(20, 0) // 0은 처음 위치, 1은 현재 위치, 2는 마지막 위치
	errCheck1(err)

	fd2 := make([]byte, 20)
	ct2, err := file.Read(fd2)
	errCheck1(err)
	fmt.Printf("읽기 작업(2) 완료 (%d bytes) (%d ret)\n\n", ct2, ol)
	fmt.Println(fd2)
	fmt.Println(string(fd2))
	fmt.Println("=======================================")

	//읽기 예제 3
	o2, err := file.Seek(0, 0)
	errCheck1(err)
	fd3 := make([]byte, 50)
	ct3, err := file.ReadAt(fd3, 8) // offset 위치부터 읽어온다.
	errCheck1(err)
	fmt.Printf("읽기 작업(3) 완료 (%d bytes) (%d ret)\n\n", ct3, o2)
	fmt.Println(fd3)
	fmt.Println(string(fd3))
	fmt.Println("=======================================")

	defer file.Close()
}
