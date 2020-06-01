package main

import (
	"fmt"
)

func start(t string) string {
	fmt.Println("start: ", t)
	return t
}

func end(t string) {
	fmt.Println("end :", t)
}

func a() {
	defer end(start("b")) //defer 함수는 바로 전 end 함수에만 적용 된다. // 중첩 함수 사용할 때 주의 사항
	fmt.Println("in a")
}

func main() {
	//example1
	a()
}
