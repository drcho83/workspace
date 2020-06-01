package main

import (
	"fmt"
)

func sum(i int, f func(int, int)) {
	f(i, 10)
}

func add(a, b int) { //a int, b int
	fmt.Println("ex1 : ", a+b)
}

func multi_value(i int) {
	i = i * 10
}

func multi_reference(i *int) {
	*i *= 10 // *i = *i*10
}

func main() {
	//함수(콜백), 참조 전달(call by reference), 값 전달(call by value)
	//example1
	sum(100, add)

	//example2
	//call by value
	a := 100
	multi_value(a)
	fmt.Println("ex2 : ", a)

	fmt.Println()
	//example3 (참조)
	b := 100
	multi_reference(&b)
	fmt.Println(b)

}
