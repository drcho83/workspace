//구조체 기본
package main

import (
	"fmt"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

//리시버를 통해서 메소드를 구조체 연결
func (a Account) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	//구조체
	//Go의 필드들의 집합체 또는 컨테이너
	//필드는 갖지만 메소드는 갖지 않는다.
	//객체지향 방식을 지원 -> 리시버를 통해 메소드랑 연결
	//상속, 객체, 클래스 개념 없음
	//구조체 -> 구조체 선언 -> 구조체 인스턴스(리시버)

	//example1
	kim := Account{number: "424-907", balance: 10000000, interest: 0.015}
	lee := Account{number: "424-908", balance: 12000000}
	Park := Account{number: "424-909", interest: 0.025}
	cho := Account{"424-909", 15000000, 0.03}

	fmt.Println("ex1:", kim)
	fmt.Println("ex1:", lee)
	fmt.Println("ex1:", Park)
	fmt.Println("ex1:", cho)
	fmt.Println()
	fmt.Println("ex2:", kim.Calculate())
	fmt.Println("ex2:", int(kim.Calculate()))
	fmt.Println("ex1:", int(lee.Calculate()))
	fmt.Println("ex1:", int(Park.Calculate()))
	fmt.Println("ex1:", int(cho.Calculate()))
	fmt.Println()
}
