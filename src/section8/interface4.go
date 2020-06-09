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

//구조체 Dog 메소드 구현

func (d Dog) run() {
	fmt.Println(d.name, "Dog is running!")
}

//Cat 메소드 구현

func (d Cat) run() {
	fmt.Println(d.name, "Cat is running!")
}

func act(animal interface{ run() }) { // 익명 인터페이스 선언
	animal.run()
}

//동물의 행동 인터페이스 선언
//duck type 메소드만으로 판단

//인터페이스를 파라메터 받는다.

func main() {
	//익명 인터페이스 사용 예제(즉시 선언 후 사용)
	//인터페이스 구현 예제

	dog := Dog{"poll", 10}
	cat := Cat{"bob", 5}

	act(dog)
	act(cat)

}
