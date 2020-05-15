//if문
package main

import "fmt"

func main() {
	//제어문(조건문)
	//IF 문: 반드시 Boolean, 검사 ->1,0 X 이렇게 안됨. true false로 명확하게 명시해 줘야 한다.
	//소괄호 사용하지 않음
	var a int = 20
	b := 20
	//예제1
	if a >= 15 {
		fmt.Println("15이상")
	}

	//예시2
	if b >= 25 {
		fmt.Println("25이상")
	}
	/*
			//에러 발생1
			if b >= 25 {
				//{ 에러 발생 반드시 조건 절에 같이 있어야 함
			}
		  //에러 발생2
		  if b >=25
		    fmt.Println("25이상")
	*/
	//에러발생3
	if c := true; c { // 초기화 하면서 조건문 진행!!!!!
		fmt.Println("True")
		if c := 40; c >= 35 {
			fmt.Println("35이상")
		}

		// 에러 발생 c +=20
	}
}
