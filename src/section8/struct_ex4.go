//상속패턴
//구조체심화
package main

import (
	"fmt"
)

type Employee struct {
	name   string
	salary float64
	bonus  float64
}

type Executives struct {
	Employee // is a 관계
	specialBonus float64
}

func (e Employee) Calulate() float64 {
	return e.salary + e.bonus
}

func main() {
	//구조체 임베디드 패턴
	//다른 관점으로 메소드를 재사용하는 장점 제공
	//상속을 허용하지 않는 Go 언어에서 메소드 상속 활용을 위한 패턴

	//example1
	ep1 := Employee{"kim", 20000000, 15000}
	ep2 := Employee{"park", 15000000, 15000}
	ex := Executives{Employee{"park", 15000000, 15000}, 1000000}

	fmt.Println("ex1 : ", int(ep1.Calulate()))
	fmt.Println("ex1 : ", int(ep2.Calulate()))
	//Employee 구조체 통해서 메소드 호출

	fmt.Println("ex1 : ", int(ex.Calulate()+ex.specialBonus))
}
