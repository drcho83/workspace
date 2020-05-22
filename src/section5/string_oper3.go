package main

import (
	"fmt"
	"strings"
)

func main() {
	//example1 결합: 일반연산
	str1 := "The Go programming language is an open source project to make programmers more productive." +
		"Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it eas" +
		"get the most out of multicore and networked machines, while its n"
	str2 := "Go is expressive, concise, clean, and efficient. Its concurrency mechanisms make it eas"

	fmt.Println("ex1 :", str1+str2)
	//example2 결합: Join

	strSet := []string{} // string 슬라이스형
	strSet = append(strSet, str1)
	strSet = append(strSet, str2)

	fmt.Println("ex2 :", strings.Join(strSet, ""))
}
