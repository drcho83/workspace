package main

import (
	"fmt"
)

func main() {
	// 부동 소수점 (실수형)
	// float32(7자리), float64(15자리)

	var num1 float32 = 0.14
	var num2 float32 = .75647
	var num3 float32 = 442.0378373
	var num4 float32 = 10.0

	//지수 표기법
	var num5 float32 = 14e6
	var num6 float32 = .156875E+3
	var num7 float64 = 5.32521e-10

	fmt.Println("ex1: ", num1)
	fmt.Println("ex1: ", num2)
	fmt.Println("ex1: ", num3)
	fmt.Println("ex1: ", num4-0.1)
	fmt.Println("ex1: ", float32(num4-0.1))
	fmt.Println("ex1: ", float64(num4-0.1))
	fmt.Println("ex1: ", num5)
	fmt.Println("ex1: ", num6)
	fmt.Println("ex1: ", num7)

}
