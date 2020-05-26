package main

import (
	"fmt"
)

func main() {
	//Slice 추가 및 병합
	//example1
	s1 := []int{1, 2, 3, 4, 5} // 5, 10 -> 7,10 -> 만약 10을 초과하면 용량이 20으로 늘어난다. // 성능 이슈 발생
	s1 = append(s1, 6, 7)

	s2 := []int{8, 9, 10, 11, 12}
	s3 := []int{13, 14, 15, 16, 17}

	s2 = append(s1, s2...) // 슬라이스에 슬라이스 삽입 시 ... 사용
	s2 = append(s2, s3[0:3]...)

	fmt.Println("ex1 :", s1)
	fmt.Println("ex1 :", s2)
	fmt.Println("ex1 :", s3)

	//example2
	s4 := make([]int, 0, 5) //길이가 0, 용량이 5인데 15번 i값이 들어가므로 길이가 15 용량이 20이 된다. 용량이 두 배씩 증가된다.
	for i := 0; i < 15; i++ {
		s4 = append(s4, i)
		fmt.Printf("ex2 -> len: %d, cap: %d, value: %v\n", len(s4), cap(s4), s4)
	}

}
