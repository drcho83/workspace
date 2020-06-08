//구조체심화
package main

import (
	"fmt"
)

type Account struct {
	number   string
	balance  float64
	interest float64
}

//생성자 리턴
func NewAccount(number string, balance float64, interest float64) *Account { // 포인터 반환 아닌 경우 값 복사
	return &Account{number, balance, interest}
}

func main() {
	//구조체 생성자 패턴 예제
	//example1
	kim := Account{number: "245-901", balance: 10000000, interest: 0.015} // 선언과 동시에 구조체 값 선언
	var lee *Account = new(Account)
	lee.number = "245-901"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex1: ", kim)
	fmt.Println("ex1: ", lee)

	park := NewAccount("4119-111", 17000000, 0.04)
	fmt.Println("ex1: ", park)
}
