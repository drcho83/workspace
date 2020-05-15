//Go의 특징
package main

import (
	"fmt"
)

func main() {
	//문자 끝 세미콜런(;) 주의
	//자동으로 끝에 세미콜론 삽입
	// 두 문장을 한 문장에 표현할 경우 명시적으로 세미콜론 사용(권장하지 않음)
	// 반복문 및 제어푼 (조건문), if / for 사용할 경우에 주의

	//example1
	for i := 0; i <= 10; i++ {
		fmt.Println("ex1: ", i)
		fmt.Println(i) //가능

		fmt.Print("ex1: ", i)

	}
	fmt.Println() //가능
	for j := 10; j >= 0; j-- {
		fmt.Println("ex2", j)
	}
}
