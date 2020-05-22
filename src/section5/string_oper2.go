package main

import (
	"fmt"
)

func main() {
	str1 := "golang"
	str2 := "world"

	fmt.Println("ex1 :", str1 == str2)
	fmt.Println("ex2 :", str1 != str2)
	fmt.Println("ex3 :", str1 > str2)
	fmt.Println("ex3 :", str1 < str2) // Go 문자열 -> 아스키 코드에 대한 사전 식 비교 //abc, ABC 비교시 대문자 ABC 가 더 크다
}
