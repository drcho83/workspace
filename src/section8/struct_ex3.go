//구조체심화
package main

import (
	"fmt"
)

//func CalulateD(a Account) {
//	a.balance = a.balance + (a.balance * a.interest)
//}

//func CalulateP(a *Account) {
//	a.balance = a.balance + (a.balance * a.interest)
//}
//리시버 형태로 선언할 경우 포인터 형식 혹은 일반 형식으로 맞춰서 할 필요가 없다. 위에 처럼..
func (a *Account) CalulateP(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a Account) CalulateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account) CalulatePP() {
	a.balance = a.balance + (a.balance * a.interest)
}

func (a Account) CalulateDD() {
	a.balance = a.balance + (a.balance * a.interest)
}

type Account struct {
	number   string
	balance  float64
	interest float64
}

func main() {
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-901", 10000000, 0.015}

	fmt.Println("ex1: ", kim)
	fmt.Println("ex1: ", lee)

	fmt.Println()

	kim.CalulateD(10000000)
	lee.CalulateP(10000000)
	///////////////////////////////
	kim.CalulateDD()
	lee.CalulatePP()
	fmt.Println("ex2: ", int(kim.balance))
	fmt.Println("ex2: ", int(lee.balance))
	fmt.Println()
	fmt.Println("ex2: ", int(kim.balance))
	fmt.Println("ex2: ", int(lee.balance))
}
