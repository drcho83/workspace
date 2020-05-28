package main

import (
	"fmt"
)

func rptc(n *int) {
	*n = 77 // 주소값을 전달할 경우 역참조 되어 번지를 찾아가서 77로 번지를 변경함
}

func vptc(n int) {
	n = 77 //복사에서 넘어 갔기 때문에 원본 값은 변경 없음, 함수 내에서만 n이 77로 변경됨
}

func main() {
	//포인터의 값 전달
	//함수, 메서드 호출 시에 매개변수 값을 복사 전달 -> 함수, 매서드 내에서는 원본 값 변경 불가능
	//원본 값 변경 위해서 포인터로 전달
	//특히 크키가 큰 배열인 경우 값 복사시 시스템 부하 -> 포인터 전달로 해결(슬라이스, 맵 참조 전)

	var a int = 10
	var b int = 10

	fmt.Println("ex1 :", a)
	fmt.Println("ex1 :", b)

	fmt.Println()

	rptc(&a)
	//rptc(a) // 에러발생
	vptc(b)
	//vptc(&b) //에러발생
	fmt.Println("ex2 :", a) //  a 값이 77로 변경됨
	fmt.Println("ex2 :", b) // 아무일도 안일어남

}
