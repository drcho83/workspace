package main

import (
	"fmt"
	"sort"
)

func main() {
	//슬라이스 추축 및 정령
	//slice[i:j] -> i 부터 j-1 까지 추출
	//slice[i:] -> i부터 마지막까지 추출
	//slice[:j] -> 처음부터 j-1끼지 추출
	//slice[:] => 처음부터 마지막까지 추출

	//example1
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("ex1: ", slice1[:])
	fmt.Println("ex1: ", slice1[0:])
	fmt.Println("ex1: ", slice1[:5])
	fmt.Println("ex1: ", slice1[0:len(slice1)]) // 0-9
	fmt.Println("ex1: ", slice1[3:])
	fmt.Println("ex1: ", slice1[0:3])
	fmt.Println("ex1: ", slice1[1:3])

	//정렬
	//sort 패키지: https://golan.org/pkg/sort
	//example2
	slice2 := []int{3, 6, 10, 9, 1, 4, 5, 8, 2, 7}
	slice3 := []string{"b", "e", "f", "a", "c"}

	fmt.Println("ex2 :", sort.IntsAreSorted(slice2)) // 정렬이 되어 있는지 frue/false 물음
	sort.Ints(slice2)                                // 실제 정렬하는 Method
	fmt.Println("ex2 :", slice2)

	fmt.Println()

	fmt.Println("ex2 :", sort.StringsAreSorted(slice3))
	sort.Strings(slice3) // 실제 정렬하는 Method
	fmt.Println("ex2 :", slice3)
}
