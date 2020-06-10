//인터페이스고급1
package main

import "fmt"

func printContents(v interface{}) {
	fmt.Printf("type: (%T)", v)
	fmt.Println("Ex1: ", v)
}

func main() {
	//인터페이스 활용(빈 인터페이스)
	//빈 인터페이스: 함수 매개변수, 리턴 값, 구조체 필드 등으로 사용 가능 -> 어떤 타입으로도 변환 가능
	//동접타입 으로 생각 하면 쉽다. (타 언어의 Object 타입)
	var a interface{}
	printContents(a)
	a = 7.5
	printContents(a)
	a = "golan"
	printContents(a)
	a = true
	printContents(a)
	a = nil
	printContents(a)
}
