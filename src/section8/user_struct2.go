package main

import (
	"fmt"
)

type cnt int // 사용자 정의 타입 int형

func main() {
	//기본 자료형 사용자 정의 타입\
	a := cnt(15)
	fmt.Println("ex1 :", a)
	var b cnt = 15
	fmt.Println("ex1 :", b)

	//example3
	//testConverT(b) 에러 발생
	testConverT(int(b)) //강제 형변환, 사용자 정의 타입 <-> 기본 타입: 매개 변수 전달 시에 변환해야 사용 가능
	testConverD(b)
}

func testConverT(i int) {
	fmt.Println("ex3 : (default type)", i)
}
func testConverD(i cnt) {
	fmt.Println("ex4 : (custom type)", i)
}
