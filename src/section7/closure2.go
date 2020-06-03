package main

import (
	"fmt"
)

func main() {
	//example1
	cnt := increaseCnt()
	fmt.Println("Ex1 : ", cnt())
	fmt.Println("Ex1 : ", cnt())
	fmt.Println("Ex1 : ", cnt())
	fmt.Println("Ex1 : ", cnt())
	fmt.Println("Ex1 : ", cnt())
}
func increaseCnt() func() int { // 함수에서 함수 리턴
	n := 0 //지역변수(캡쳐됨)
	return func() int {
		n += 1
		return n
	}
}
