package main

import (
	"fmt"
)

func main() {
	//example1
	var str1 string = "Golang"
	var str2 string = "World"

	fmt.Println("ex1: ", str1[0:2], str1[0]) //slice 문자 출력, 인덱스 출력 하면 해당 코드 값 출력!!!!!
	fmt.Println("ex1: ", str2[3:], str1[0])
	fmt.Println("ex1: ", str2[:4])
	fmt.Println("ex1: ", str1[1:3])
}
