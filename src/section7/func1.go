package main

import (
	"fmt"
	"strconv"
)

//함수 선언 위치는 어느 곳이든 가능
func hellogolang() {
	fmt.Println("ex1: Hello Golang!!")
}
func say_one(m string) {
	fmt.Println("ex2 :", m)
}

func sum(x int, y int) int {
	return x + y
}
func main() {
	//선언: func 키워드로 선언
	//func 함수명(매개변수)(반환타입 or 반환 값 변수명): 반환 값 존재
	//func 함수명()(반환타입 or 반환 값 변수명): 매개변수 없으나 리턴값 존재
	//func 함수명(매개변수): 매개변수 존재, 반환 값 없음
	//func 함수명():매개변수 없음, 반환 값 없음
	//타 언어와 달리 반환 값 (return value) 다수 가능

	//example1
	hellogolang()
	//example2
	say_one("Hello World!!")
	//example3
	result := sum(5, 5)
	fmt.Println(result)
	fmt.Println(sum(5, 5))
	//숫자를 문자열로 변환
	fmt.Println("ex3 : ", strconv.Itoa(sum(5, 5))) // int를 문자열로 변환

}
