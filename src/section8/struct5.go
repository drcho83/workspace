//익명 구조체
package main

import (
	"fmt"
	"reflect"
)

type Car struct { // 대문자 시작 시 패키지 외부에서 참조 가능
	name    string "차량명"
	color   string "색상"
	company string "제조사"
}

func main() {
	//구조체 익명 선언 및 활용
	//example1 필드테크 사용 예제
	tag := reflect.TypeOf(Car{})
	for i := 0; i < tag.NumField(); i++ {
		fmt.Println("ex1: ", tag.Field(i).Tag, tag.Field(i).Name, tag.Field(i).Type)
	}
}
