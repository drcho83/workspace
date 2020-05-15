//변수1
package main

import "fmt"

func main() {
	//기본 초기화
	//정수타입: 0, 실수(소수점),문자열,boolean
	//변수형: 숫자 첫클자x, 대소문자구분 O, 문자, 숫자, 밑줄, 특수기호 사용 가능
	//변수 및 상수 사유: 함수 내외 사용 하기 위해서
	var a int
	var b string
	var c, d, e int
	var f, g, h int = 1, 2, 3
	var i float32 = 11.4
	var j = "hi! gorang!!"
	var k = 4.75 // 선언 동시 초기화
	var l = "h1 seoul!!"
	var m = true

	a = 4
	b = "aaaa"
	e = 33

	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)
	fmt.Println("e: ", e)
	fmt.Println("f: ", f)
	fmt.Println("g: ", g)
	fmt.Println("h: ", h)
	fmt.Println("i: ", i)
	fmt.Println("j: ", j)
	fmt.Println("k: ", k)
	fmt.Println("l: ", l)
	fmt.Println("m: ", m)
}
