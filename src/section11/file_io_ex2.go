package main

import (
	"bufio"
	"fmt"
	"os"
)

func errCheck(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	//파일 읽기, 버퍼 사용 쓰기 -> bufio 패키지 활용
	//ioutil, bufio 등은 io.Reader, io.Writer 인터페이스를 구현 함 -> 즉, writer, read 메소드를 사용 가능
	/*
		type Reader interface {
		Read(p []byte){n int, err error}
		}
		type Writer interface {
		Write(p []byte){n int, err error}
		}
	*/

	//즉, bufio의 newReader, newWriter를 통해 객체 변환 후 메소드 사용
	//bufio(Bufferd io) 패키지
	//https://golang.org/pkg/bufio

	// 파일 열기
	//두 번째 매개 변수 확인
	//https://golang.org/pkg/os/#pkg-constants

	/*
		설명
		a -----> a
		b -----> ab
		c -----> abc
		d -----> abcd
		e -----> e    -----> abcd
		f -----> ef   -----> abcd
		g -----> efg  -----> abcd
		h -----> efgh -----> abcd
		i -----> i		-----> abcdefgh
	*/
	file, err := os.OpenFile("test_write2.txt", os.O_CREATE|os.O_RDWR, os.FileMode(0777))
	errCheck(err)
	//Bufio 파일 쓰기 예제
	wt := bufio.NewWriter(file) //writer 반환 (버퍼를 사용)
	wt.WriteString("Hello Golang \n FIle write Test1\n")
	wt.Write([]byte("Hello Golang \n FIle write Test2\n"))

	fmt.Printf("사용한 버퍼 사이즈: (%d bytes)\n ", wt.Buffered())
	fmt.Printf("남은 버퍼 사이즈: (%d bytes)\n ", wt.Available())
	fmt.Printf("전체 버퍼 사이즈: (%d bytes)\n ", wt.Size())
	wt.Flush() // 해주지 않으면 생성되지 않음
	fmt.Println("쓰기 작업 완료\n")
	fmt.Println("===============")

	//읽기 작업 완료
	rt := bufio.NewReader(file)
	fi, err := file.Stat()
	errCheck(err)

	b := make([]byte, fi.Size())
	fmt.Println("파일 정보 출력: ", fi)
	fmt.Println("파일 이름: ", fi.Name())
	fmt.Println("파일 크기: ", fi.Size())
	fmt.Println("파일 수정 시간: ", fi.ModTime())

	fmt.Println("===============================================")
	file.Seek(0, os.SEEK_SET)

	data, _ := rt.Read(b) // 읽기 (Readline, ReadByte, ReadBytes 등)
	//rt.Read(b)

	fmt.Printf("전체 버퍼 사이즈: (%d bytes)\n", rt.Size())
	fmt.Printf("읽기 작업 완료:(%d bytes)\n", data)

	fmt.Println("===============================================")
	fmt.Println(string(b))
	fmt.Println("===============================================")
	defer file.Close()

}
