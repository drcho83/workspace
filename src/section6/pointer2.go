package main

import (
	"fmt"
)

func main() {
	//example1
	i := 7
	p := &i
	fmt.Println("ex1 :", i)
	fmt.Println("ex1 :", *p) //참조값만 전달 하였는데 역참조하여 i값을 보여
	fmt.Println("ex1 :", &i)
	fmt.Println("ex1 :", p)
	*p++ //1증가
	fmt.Println()

	fmt.Println("ex1 :", i)
	fmt.Println("ex1 :", *p) //참조값만 전달 하였는데 역참조하여 i값을 보여
	fmt.Println("ex1 :", &i)
	fmt.Println("ex1 :", p)
	fmt.Println()
	*p = 7777 // 포인터 변수 역 참조값 변경
	fmt.Println("ex1 :", i)
	fmt.Println("ex1 :", *p) //참조값만 전달 하였는데 역참조하여 i값을 보여
	fmt.Println("ex1 :", &i)
	fmt.Println("ex1 :", p)

	i = 77
	fmt.Println()
	fmt.Println("ex1 :", i)
	fmt.Println("ex1 :", *p) //참조값만 전달 하였는데 역참조하여 i값을 보여
	fmt.Println("ex1 :", &i)
	fmt.Println("ex1 :", p)
}
