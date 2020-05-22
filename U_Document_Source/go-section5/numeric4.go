//데이터 타입 : Numeric(4)
package main

import "fmt"

func main() {
	//데이터 타입 : 숫자형
	//복소수 형(complex number)
	//complex64(32bit 실수 + 허수)
	//complex128(64bit 실수 + 허수)

	//예제1
	var num1 complex64 = 5 + 7i
	num2 := 8 + 1i
	num3 := complex(3, 2) //complex128
	var num4 complex128 = 9 + 3i
	num5 := complex64(2 + 3i)

	fmt.Println("ex1 : ", num1)
	fmt.Println("ex1 : ", num2)
	fmt.Println("ex1 : ", num3)
	fmt.Println("ex1 : ", num4)
	fmt.Println("ex1 : ", num5)

	//예제2
	//real() : 실수부 출력
	//imag() : 허수부 출력
	fmt.Println("ex2 : ", num1, real(num1), imag(num1))
	fmt.Println("ex2 : ", num2, real(num2), imag(num2))
	fmt.Println("ex2 : ", num3, real(num3), imag(num3))
	fmt.Println("ex2 : ", num4, real(num4), imag(num4))
	fmt.Println("ex2 : ", num5, real(num5), imag(num5))

}
