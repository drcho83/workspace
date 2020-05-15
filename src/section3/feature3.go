package main

import (
	"fmt"
)

func main() {
	//코드 서식 지원하는 유틸
	//여러 사람이 코딩한 것을 한 사람이 코딩 한 것 같은 일관성 유지 => fmt
	// 코드 스타일 유지
	//gofmt -h: 사용법
	//gofmt -w : 원본파일에 반영

	//example1
	for i := 0; i <= 100; i++ {
		fmt.Println("ex1:", i)
	} // fmt 가 알아서 변경해 준다.
}
