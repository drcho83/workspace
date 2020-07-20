package main

import (
	"fmt"
	"log"
)

func main() {
	//보통 컴파일 언어의 경우 컴파일 시 에러 확인 가능
	//동적 타입의 경우 컴파일 과정이 없음
	//사용자가 에러 발생 가능
	//Panic 함수는 호출 즉시, 해당 메소드를 즉지 중지시키고  defer 함수를 호출하고 자기자신을 호출한 곳으로 리턴
	//런타임 이외에 사용자가 코드 흐름에 따라 에러를 발생 시킬 때 중요!!
	//문법적인 에러는 아니지만, 논리적인 코드 후름에 따른 에러 발생 처리 가능
	fmt.Println("Start Main")
	//panic("Error occured: user stopped!!") // 방법1
	log.Panic("Error occured: user stopped!!") // 방법2
	fmt.Println("End Main")                    // 실행 불가!!

}
