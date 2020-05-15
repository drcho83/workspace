//Go의 특징
package main

import (
	"fmt"
)

func main() {
	//go: 모호하거나, 애매한 문법 금지
	// 후위,후치 연산자 허용 i++, 전치(전위) 연산자 비허용 (++i)는 사용 불가
	//증감연산 반환 값 없음
	//포인터(Pointer 허용, 연산 비허용)
	//주석 (//,/**/) 사용법 숙지
	sum, i := 0, 0
	for i <= 100 {
		//sum += i++ 에러 발생
		sum += i
		i++
		//++i 에러 발생
	}
	fmt.Println("ex1", sum)
}
