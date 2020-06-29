//함수 심화 / 익명함수 /사용 후 소멸
package main

import (
	"fmt"
)

func main() {
	//즉시 실행 (선언과 동시에)

	//익명 example1
	func() {
		fmt.Println("ex1: anonymous func!!")
	}()

	//example2
	msg := "hello golang!"

	func(m string) {
		fmt.Println("ex2 :", m)
	}(msg)

	//example3
	func(x, y int) {
		fmt.Println("ex3 :", x+y)
	}(10, 10)

	//example4
	r := func(x, y int) int {
		return x * y
	}(10, 100)
	fmt.Println("ex4 :", r)
}
