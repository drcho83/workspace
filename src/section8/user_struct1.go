//사용자 정의 타입(1)
package main

import "fmt"

type Car struct {
	name  string
	color string
	price int64
	tax   int64
}

//구조체 <-> 매소드 바인딩

func Price(c Car) int64 {
	return c.price + c.tax
}

func main() {
	//GO -> 객체 지향 타입을 구조체로 정의한다.(클래스, 상속 개념 없음)
	//객체 지향 -> 클래스 (속성: 멤버변수 , 기능(상태:메소드)): 코드의 재사용성, 코드의 관리가 용이, 신뢰성이 높은 프로그램
	//객체 지향 언어일까? 맞음
	//GO는 전형적인 객체지향의 특징을 가지고 있지 않지만, 인터페이스 -> 다형성 지원, 구조체를 통한 클래스형태의 코딩 가능
	//객체지향의 기본 개념 -> GO에서 포함 하고 있다. -> 객체 지향 프로그램밍 언어
	//상태, 메소드 분리해서 정의(결합성 없음)
	//사용자 정의 타입: 구조체, 인터페이스, 기본 타입(int, floot,string...), 함수
	//구조체와 -> 메소드 연결을 통해서 타 언어의 클래스 형식 처럼 사용 가능(객체 지향)
	//example1
	bmw := Car{name: "520d", price: 50000000, color: "white", tax: 500000}
	benz := Car{name: "520d", price: 60000000, color: "white", tax: 600000}

	fmt.Println("ex1: ", bmw, &bmw) //&탐조타입//원본이 변경 됨
	fmt.Println("ex1: ", benz, &benz)

	fmt.Println("ex2: ", Price(bmw))
	fmt.Println("ex2: ", Price(benz))

}
