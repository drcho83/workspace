//데이터 타입 : Bool(1)
package main

import "fmt"

func main() {
	//Bool
	//조건부 논리 연산자랑 주로 사용 : !, || , &&
	//암묵적 형변환 X : 0, nil -> false 변환 없음

	//예제1
	var b1 bool = true
	var b2 bool = false
	b3 := true
	//b4 := 1

	fmt.Println("ex1 : ", b1)
	fmt.Println("ex1 : ", b2)
	fmt.Println("ex1 : ", b3 == b3)
	fmt.Println("ex1 : ", b1 && b3)
	fmt.Println("ex1 : ", b1 || b2)
	fmt.Println("ex1 : ", !b2)

	/*
		예외 발생
		if b4 {
			fmt.Println("true")
		}
	*/

}
