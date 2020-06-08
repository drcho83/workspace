//익명 구조체
package main

import (
	"fmt"
)

type Car struct { // 대문자 시작 시 패키지 외부에서 참조 가능
	name    string "차량명"
	color   string "색상"
	company string "제조사"
	detail  spec   "상세" //구조체 안에 구조체
}

type spec struct { //소문자
	length int "잔장"
	height int "전고"
	width  int "전축"
}

func main() {
	//중첩 구조체
	ca1 := Car{
		"250d",
		"siver",
		"bmw",
		spec{4000, 1000, 2000},
	}
	//내부 spec 구조체 값 출력
	fmt.Println("ex1: ", ca1.name)
	fmt.Println("ex1: ", ca1.color)
	fmt.Println("ex1: ", ca1.company)
	fmt.Printf("ex1: %#v\n", ca1.detail)
	fmt.Printf("ex1: %#v\n", ca1.detail.length)

	//example2
	//내부 spec 구조체 필드 값 출력
	fmt.Println("ex1: ", ca1.detail.height)
	fmt.Println("ex1: ", ca1.detail.length)
	fmt.Println("ex1: ", ca1.detail.width)
}
