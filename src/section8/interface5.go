package main

import (
	"fmt"
)

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

func printValue(s interface{}) {
	fmt.Println("ex1 :", s)
}

func main() {
	//인터페이스 활용 (빈 인터페이스)
	//함수내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다.(만능) -> 모든 타입 지정가
	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	printValue(dog)
	printValue(cat)
	printValue(15)   // 어떠한 자료형도 모두 받는다
	printValue(true) // 어떠한 자료형도 모두 받는다
	printValue(25.5) // 어떠한 자료형도 모두 받는다
	printValue("animal")
	printValue([]Dog{})
	printValue([5]Dog{})
}
