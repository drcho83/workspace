package main

import (
	"fmt"
)

func stack() { // fisrt in first out
	for i := 1; i <= 10; i++ {
		defer fmt.Println("ex1 :", i) // fmt 자체도 함수이기 때문 defer를 사용할 수 있따.!!
	}
}
func main() {
	//example1
	stack()
}
