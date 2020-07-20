package main

import (
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
	// 파일 쓰기
	//create : 새 파일 작성 및 파일 열기
	//close: 리소스 닫기
	//write <- byte 불편,writeString, writeAt: 파일 쓰기
	//각 운영체제 권한 주위(오류 메시지 확인)
	//예외 처리 정말 중요!!
	//https://golang.org/pkg/os

	//example
	file, err := os.Create("test_write.txt")
	errCheck1(err)

	defer file.Close()

	//write example1
	s1 := []byte{115, 111, 109, 101, 115}
	n1, err := file.Write([]byte(s1)) // 문자열 -> byte 슬라이스 형으로 변환 후 쓰
	errCheck2(err)

	fmt.Printf("쓰기 작업(1) 완료 (%d bytes \n)", n1)
	file.Sync() // Write Commit 하는 행위(Stable)

	//example2
	s2 := "\nHello Golan\n file write test! -2 \n"
	n2, err := file.WriteString(s2)
	errCheck2(err)
	fmt.Printf("쓰기 작업(2) 완료 (%d bytes \n)", n2)
	file.Sync() // Write Commit 하는 행위(Stable)

	s3 := "Test WriteAt! -3\n"
	n3, err := file.WriteAt([]byte(s3), 100) //space bar 를 70번 누른 위치에서 시작? //len(offset) 조절하면서 테스
	errCheck1(err)
	fmt.Printf("쓰기 작업(3) 완료 (%d bytes \n)", n3)
	file.Sync() // Write Commit 하는 행위(Stable)
	//example4
	n4, err := file.WriteString("Hello Golang! \n File Write Test! - 4")
	errCheck1(err)
	fmt.Printf("쓰기 작업(4) 완료 (%d bytes \n)", n4)
	file.Sync() // Write Commit 하는 행위(Stable)

}
