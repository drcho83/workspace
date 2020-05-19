package main

import (
	"fmt"
)

func main() {
	//example1
	var n1 uint8 = 125
	var n2 uint8 = 90

	fmt.Println("ex1: ", n1+n2)
	fmt.Println("ex1: ", n1-n2)
	fmt.Println("ex1: ", n1*n2)
	fmt.Println("ex1: ", n1/n2)
	fmt.Println("ex1: ", n1%n2)
	fmt.Println("ex1: ", n1<<2)
	fmt.Println("ex1: ", n1>>2)
	fmt.Println("ex1: ", ^n1)

	//example2
	var n3 int = 12
	var n4 float32 = 8.6
	var n5 uint16 = 1024
	var n6 uint32 = 120000

	fmt.Println("ex3: ", float32(n3)+n4) //12.0 로 형변환 되어 계산됨
	fmt.Println("ex3: ", n3+int(n4))     // 8.2 에서 소숫점 아래가 버려서 계산됨 --> 형 변환이 중요한 사례
	fmt.Println("ex3: ", n5+uint16(n6))  //16으로 형변환 되면서 데이터를 표한할 수 있는 범위가 넘어가 강제적으로 1024로 설정 되어 값이 계산 되었음 --> 데이터 손실
}
