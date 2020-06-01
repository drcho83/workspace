package main

import "fmt"

func main() {
	//포인터
	//Go: 포인터 지원(c)
	//변수의 지역성, 연속된 메모리의 참조..., 힙, 스텍....
	//파이썬, 자바 (JRE) -> 컴파일러, 인터프리터
	//포인터를 지원하지 않는 언어: C#, JAVA 등
	//Golang: 주소의 값은 집적 변경 불가능(잘못된 코딩으로 인한 버그 방지)
	//*(에스터리스크) 로 사용
	//nil 로 초기화(nil != 0)

	//example1
	var a *int            // 방법1
	var b *int = new(int) // 방법2 -> 정석

	fmt.Println(a) //&
	fmt.Println(b)

	i := 7
	fmt.Println("ex1 :", i, &i)

	fmt.Println()
	a = &i // &주소값 전달
	b = &i

	fmt.Println("ex1: ", a, &a, *a, &i) //&a 자체의 주소값 // *a 역참조라 한다.
	fmt.Println("ex1: ", b, &b, *b, &i)

	fmt.Println()

	var c = &i
	d := &i

	*d = 10
	*a = 77
	*b = 99

	fmt.Println("ex1: ", c, &c, *c, &i) //&a 자체의 주소값 // *a 역참조라 한다.
	fmt.Println("ex1: ", d, &d, *d, &i)
	fmt.Println("ex1: ", b, &b, *b, &i)

}
