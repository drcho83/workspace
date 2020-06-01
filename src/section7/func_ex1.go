//함수 심화
package main

import (
	"fmt"
)

func multiply(n ...int) int { // int 형으로 가변으로 받는다.
	tot := 1

	for _, value := range n {
		//		fmt.Println(s, value)
		tot *= value

	}
	return tot
}

func sum(n ...int) int {
	tot := 0

	for _, value := range n {
		//		fmt.Println(s, value)
		tot += value

	}
	return tot
}

func prtWorld(msg ...string) {
	for _, value := range msg {
		fmt.Println("ex2 :", value)
	}
}

func main() {
	//함수 고급
	//가변 인자 실습 (매개 변수 개수가 동적으로 변할 때 - 정해져 있지 않음)

	x := multiply(5, 6, 7, 8, 9, 10, 11) // 너어 갈 때 ky, value 전달됨
	y := sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println("ex1: ", x)
	fmt.Println("ex2: ", y)
	prtWorld("a", "apple", "test", "golang")

	//example3
	a := []int{1, 2, 3, 4, 5}
	m := multiply(a...)
	n := sum(a...)
	fmt.Println("ex2: ", m)
	fmt.Println("ex2: ", n)
}
