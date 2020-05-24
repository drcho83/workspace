package main

import (
	"fmt"
)

func main() {
	//슬라이스(참조 타입 증명)
	//example1
	arr1 := [3]int{1, 2, 3}
	var arr2 [3]int

	arr2 = arr1
	arr2[0] = 7

	fmt.Println("ex1: ", arr1)
	fmt.Println("ex1: ", arr2)
	fmt.Println()

	//example2(slice)
	slice1 := []int{1, 2, 3}
	var slice2 []int

	slice2 = slice1
	slice2[0] = 7

	fmt.Println("ex2: ", slice1)
	fmt.Println("ex2: ", slice2)
	fmt.Println()

	//example3 (슬라이스 예외 상황)
	slice3 := make([]int, 50, 100) //50까지 길이만큼만 0으로 초기화 됨
	fmt.Println("ex3: ", slice3[4])
	//fmt.Println("ex3: ", slice3[50]) // error 발생 용량 만큼 0으로 초기화 되는 것이 아니기 때문에 인덱스에러 발생

	//example4
	for i, v := range slice1 {
		fmt.Println("ex4: ", i, v)
	}
}
