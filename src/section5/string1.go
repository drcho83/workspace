//데이터 타입1: 문자열(1)
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//문자열
	//큰 따옴표"", 백스쿼드``
	//golang: 문자의 char 타입을 제공하지 않음 -> rune(int32) 로 문자 코드 값으로 표현
	//문자: ''로 작성
	//자주 사용하는 escape:\\,\',\" \이후에 나오는 문자를 그대로 표현
	//\a(콘솔벨), \b(스페이스),\f(쪽바꿈),\n(줄바꿈),\r(복귀),\t(탭)

	//example1
	var str1 string = "c:\\go_study\\src\\" //->  c:\go_study\src\ 로 변환 됨
	str2 := `c:\go_study\src\`              //<- escape 문자를 사용하지 않더라도 그대로 사용할 수 있어 더 간편함 위의 st1 과 결과는 같음

	fmt.Println("ex1: ", str1)
	fmt.Println("ex1: ", str2)

	//example2
	var str3 string = "Hello, World!!" //1바이트
	var str4 string = "안녕하세요."         //3바이트
	var str5 string = "\ud55c\uae00"   //사용하지 않음 "한글"

	fmt.Println()
	fmt.Println("ex2 :", str3)
	fmt.Println("ex2 :", str4)
	fmt.Println("ex2 :", str5)

	//example3
	//길이 (바이트 수)
	fmt.Println("ex3 :", len(str3))
	fmt.Println("ex3 :", len(str4))
	fmt.Println("ex3 :", len(str5))

	//example4
	//길이(실제길이)
	fmt.Println("ex4 :", utf8.RuneCountInString(str3))
	fmt.Println("ex4 :", utf8.RuneCountInString(str4))
	fmt.Println("ex4 :", utf8.RuneCountInString(str5))
	fmt.Println("ex4: ", len([]rune(str3)))
	fmt.Println("ex4: ", len([]rune(str4)))
	fmt.Println("ex4: ", len([]rune(str5)))
}
