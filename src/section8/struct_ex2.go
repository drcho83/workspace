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

func CalulateD(a Account) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalulateP(a *Account) {
	a.balance = a.balance + (a.balance * a.interest)
}
func main() {
	kim := Account{"245-901", 10000000, 0.015}
	lee := Account{"245-901", 10000000, 0.015}

	fmt.Println("ex1: ", kim)
	fmt.Println("ex1: ", lee)
	CalulateD(kim)
	CalulateP(&lee)
	fmt.Println("ex2: ", int(kim.balance))
	fmt.Println("ex2: ", int(lee.balance))
}
