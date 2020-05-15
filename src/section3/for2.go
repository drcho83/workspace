package main

import (
	"fmt"
)

func main() {
	//example1
	sum1 := 0
	for i := 0; i < 100; i++ {
		sum1 += 1
	}
	fmt.Println("ex1: ", sum1)

	//example2
	sum2, i := 0, 0
	for i <= 100 {
		sum2 += i
		i++
		//j:=i++ Go에서 후치 연산은 반환 값이 없기 때문에 바로 에러 발생 --> 단독으로 써야 함
	}
	fmt.Println("ex2 :", sum2)

	//example3
	sum3, i := 0, 0
	for { //while 형태와 비슷
		if i > 100 {
			break
		}
		sum3 += i
		i++
	}
	fmt.Println("ex3: ", sum3)

	//example4
	for i, j := 0, 0; i <= 10; i, j = i+1, j+10 {
		fmt.Println("ex4:", i, j)
	}

	/*
		//에러 발생
		for i, j := 0, 0; i <= 10; i++,j+=10 { //후치 연산은 반환 값이 없기 때문에 에러가 발생된다.
			fmt.Println("ex4:", i, j)
		}
	*/
}
