//데이터 타입 : 문자열 연산(2)
package main

import "fmt"

func main() {

	//예제1(비교)
	str1 := "Golang"
	str2 := "World"

	fmt.Println("ex1 : ", str1 == str2) //바이트로 비교
	fmt.Println("ex1 : ", str1 != str2)
	fmt.Println("ex1 : ", str1 > str2)
	fmt.Println("ex1 : ", str1 < str2)

}
