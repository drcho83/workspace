//인터페이스고급3
package main

import (
	"fmt"
)

func checkType(arg interface{}) {
	//arg.(type)을 통해서 현재 데이터형 반환
	switch arg.(type) {
	case bool:
		fmt.Println("this is a bool", arg)
	case int, int8, int16, int32, int64:
		fmt.Println("this is a int", arg)
	case float64:
		fmt.Println("this is a float", arg)
	case string:
		fmt.Println("this is a string", arg)
	case nil:
		fmt.Println("this is a nil", arg)
	default:
		fmt.Println("what is this type?", arg)
	} // golang에서는 Break 문이 없음
}

func main() {
	//실제 타입 검사 switch 사용
	// 빈 인터페이스는 어떠한 자료형도 전달 받을 수 있으므로, 타입 체크 후 형 변화하여 사용 가능

	//example1
	checkType(true)
	checkType(1)
	checkType(22.542)
	checkType(nil)
	checkType("hello golang!!")
}
