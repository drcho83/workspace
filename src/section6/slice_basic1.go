package main

import (
	"fmt"
)

func main() {
	//slice
	//가변 -> 동적으로 크기가 늘어난다, 참조 타입
	//슬라이스 (길이 & 용량)
	//배열 vs 슬라이스 차이점 중요
	//길이고정 vs 길이가정
	//값 타입 vs 참조 타입
	//복사 전달 vs 참조 값 전달
	//전체 비교연산자 vs 비교 연산자 사용 불가능
	//대부분 슬라이스 사용한다.

	// 2가지 선언 방법 1. 배열 처럼 선언, 2. make 함수 : make(자료형, 길이, 용량(생략시 길이))

	//example1
	var slice1 []int
	slice2 := []int{}
	slice3 := []int{1, 2, 3, 4, 5}
	slice4 := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}
	//	slice2[0] = 1  // 길이를 모르기 때문에 인덱스 범위 오류 발생
	slice3[4] = 10 // 값 수정 가능

	fmt.Printf("%-5T %d %d %v\n", slice1, len(slice1), cap(slice1), slice1)
	fmt.Printf("%-5T %d %d %v\n", slice2, len(slice2), cap(slice2), slice2)
	fmt.Printf("%-5T %d %d %v\n", slice3, len(slice3), cap(slice3), slice3)
	fmt.Printf("%-5T %d %d %v\n", slice4, len(slice4), cap(slice4), slice4)
	fmt.Println()
	//example2
	var slice5 []int = make([]int, 5, 10) //길이는 5이지만 메모리 용량은 10으로 미리 확보
	// 데이터 11번째를 넣으면 메모리 자동 재할당함

	var slice6 = make([]int, 5)
	slice7 := make([]int, 5, 100)
	slice8 := make([]int, 5)
	var slice9 []int //nil 슬라이스(길이와 용량이 0)
	if slice9 == nil {
		fmt.Println("this is nil slice")
	}

	slice6[2] = 7

	fmt.Printf("%-5T %d %d %v\n", slice5, len(slice5), cap(slice5), slice5)
	fmt.Printf("%-5T %d %d %v\n", slice6, len(slice6), cap(slice6), slice6)
	fmt.Printf("%-5T %d %d %v\n", slice7, len(slice7), cap(slice7), slice7)
	fmt.Printf("%-5T %d %d %v\n", slice8, len(slice8), cap(slice8), slice8)
	fmt.Printf("%-5T %d %d %v\n", slice9, len(slice9), cap(slice9), slice9)

	//make 사용 시 0으로 초기화 까지 진행됨
}
