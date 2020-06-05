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
func (a Account) Calculate() float64 { //func (a *Account) Calculate() float64 { --> 이것도 가능
	return a.balance + (a.balance * a.interest)
}

func main() {
	//다양한 선언 방법
	//&struct, struct : &struct 포인터를 받아오고 역참조를 또 하기 때문에 속도는 조금 느리다.
	//인터페이스 메소드를 선언만 한 후 -> 오버라이딩 해서 메소드에 포인터 리시버를 사용할 경우 받드시 &struct로 넘겨야 한다.
	//선언 방법1
	var kim *Account = new(Account)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	//선언 방버2
	hong := &Account{number: "245-901", balance: 10000000, interest: 0.015}

	//선언 방법3
	lee := new(Account) // new 로 할 경우 선언만 할 수 있고, 동시에 값을 넣을 수는 없다
	lee.number = "245-901"
	lee.balance = 10000000
	lee.interest = 0.015

	fmt.Println("ex1 :", kim)
	fmt.Println("ex1 :", hong)
	fmt.Println("ex1 :", lee)
	fmt.Println()
	fmt.Printf("ex2 :%#v\n", kim)
	fmt.Printf("ex2 :%#v\n", hong)
	fmt.Printf("ex2 :%#v\n", lee)
	fmt.Println()
	fmt.Println("ex1 :", int(kim.Calculate()))
	fmt.Println("ex1 :", int(hong.Calculate()))
	fmt.Println("ex1 :", int(lee.Calculate()))

}
