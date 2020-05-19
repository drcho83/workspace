package main

import (
	"fmt"
)

func main() {
	//데이터 타입: 숫자형
	//정수형 문자 출력
	//example1
	//아스키(영문)

	var char1 byte = 72   //unit8 동일 10진수
	var char2 byte = 0110 //unit8 동일 8진수
	var char3 byte = 0x48 //unit8 동일 16진수

	//example2
	//unicode(한글)
	var char4 rune = 50556   //유니코드 10진수
	var char5 rune = 0142574 //44032 8진수
	var char6 rune = 0xc57c  //44032 16진수

	fmt.Printf("%c %c %c\n", char1, char2, char3) //%c 문자 10진
	fmt.Printf("%d %d %d\n", char1, char2, char3) //%d digit 10진수
	fmt.Printf("%ㅇ %o %x\n", char1, char2, char3) //%o 8진수, %x 16진수

	fmt.Printf("%c %c %c\n", char4, char5, char6) //%c 문자
	fmt.Printf("%d %d %d\n", char4, char5, char6) //%d digit
	fmt.Printf("%ㅇ %o %x\n", char4, char5, char6)
}
