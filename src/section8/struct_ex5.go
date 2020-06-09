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
	Employee
	specialBonus float64
}

func (e Employee) Calulate() float64 {
	return e.salary + e.bonus
}

func (e Executives) Calulate() float64 {
	return e.salary + e.bonus + e.specialBonus
}

func main() {
	//구조체 임베디드 메소드 오버라이딩 패턴
	//example1
	ep1 := Employee{"kim", 2000000, 150000}
	ep2 := Employee{"park", 1500000, 200000}
	ex := Executives{Employee{"park", 5000000, 1000000}, 1000000}

	fmt.Println("ex1 : ", int(ep1.Calulate()))
	fmt.Println("ex1 : ", int(ep2.Calulate()))
	//Employee 구조체 통해서 메소드 호출

	fmt.Println("ex1 : ", int(ex.Calulate()))                          // 오버라이딩 됨
	fmt.Println("ex1 : ", int(ex.Employee.Calulate()+ex.specialBonus)) // 따로 뽑아서 실행
	//fmt.Println("ex1 : ", int(ex.Calulate()+ex.specialBonus)) 오버라이딩 잘 못 된값 반환
}
