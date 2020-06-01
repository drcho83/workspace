//함수 심화
package main

import (
	"fmt"
)

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func prtHello(n int) {
	if n == 0 {
		return
	}
	fmt.Println("ex2: (", n, ")", "hi!!")
	prtHello(n - 1)
}

func main() {
	//재귀함수
	//프로그램이 보기 쉽고, 코드 간결, 오류 수정이 용이: 장점
	// 코드 이해하기 어렵고, 기억공간을 많이 사용, 무한루프 가능성

	//example1
	x := fact(5)
	fmt.Println("ex1 :", x)

	//example2
	prtHello(10)

}
