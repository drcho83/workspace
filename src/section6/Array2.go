package main

import (
	"fmt"
)

func main() {
	//배열 순회
	//example1
	arr1 := [5]int{1, 10, 100, 1000, 10000}
	//len 길이 반속
	for i := 0; i < len(arr1); i++ {
		fmt.Println("ex1: ", arr1[i])
	}

	//
	fmt.Println()

	//example2 //가장 많이 사용
	arr2 := [5]int{7, 77, 777, 7777, 77777}

	for i, v := range arr2 {
		fmt.Println("ex1: ", i, v)
	}
	//인덱스 생략
	for _, v := range arr2 {
		fmt.Println("ex1: ", v)
	}

	//인덱스 생략2
	fmt.Println()
	for v := range arr2 {
		fmt.Println("ex4 :", v)
	}

}
