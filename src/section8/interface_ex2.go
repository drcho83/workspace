//인터페이스고급2
package main

import (
	"fmt"
	"reflect"
)

func main() {
	//타입 변환
	//Type Assertion
	//실행(런타임) 시에는 인터페이스에 할당한 변수는 실제 타입으로 변환 후 사용해야 하는 경우
	//인터페이스.(타입) -> 형 변환
	//interfaceVal.(type)

	//example1
	var a interface{} = 15
	b := a
	c := a.(int)
	//d := a.(float64) // error 발생

	fmt.Println("ex1 : ", a)
	fmt.Println("ex1 : ", reflect.TypeOf(a))
	fmt.Println("ex1 : ", b)
	fmt.Println("ex1 : ", reflect.TypeOf(b))
	fmt.Println("ex1 : ", c)
	fmt.Println("ex1 : ", reflect.TypeOf(c))
	//fmt.Println("ex1 : ", d)
	fmt.Println()
	//example2 (저장된 실제 타입 검사)
	if v, ok := a.(int); ok {
		fmt.Println("ex2 : ", v, ok)
	}

}
