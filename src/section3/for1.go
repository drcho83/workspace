package main

import (
	"fmt"
)

func main() {
	//반복문 - for
	//Go에서 유일하게 제공되는 반복문
	//다양한 사용범 숙지

	//example1
	for i := 0; i < 5; i++ {
		fmt.Println("ex1: ", i)
	}
	//에러 발생1
	/*
	  for i :=0; i<5;i++
	  {

	  }
	  //에러 발새애2
	  for i :=0,i<5; i++
	  	fmt.Println("ex1: ", i)
	*/

	//example2(무한 루프)
	/*
	  for{
	      fmt.Println("ex1: Hello Golang")
	      fmt.Println("ex2: Infinite loop!")
	  }
	*/
	//example3 (Range 용법)
	loc := []string{"seoul", "busang", "incheon"}

	for index, name := range loc {
		fmt.Println("ex3: ", index, name)
	}

	for _, name := range loc {
		fmt.Println("ex3: ", name)
	}

}
