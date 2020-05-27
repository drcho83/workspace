package main

import (
	"fmt"
)

func main() {
	//map
	//맵 조회 할 경우에 주의 할 점

	//example1
	map1 := map[string]int{ //int : 0, string: "", float:0.0) <- 자동 초기화 값
		"apple":  15,
		"banana": 115,
		"orange": 1115,
		"lemon":  0,
	}
	value1 := map1["lemon"]      // value1 이 int 형으로 자동 초기화 됨
	value2, ok1 := map1["lemon"] // value2 도 map1에서 정의한 int 형으로 자동 초기화 됨
	value3, ok := map1["kiwi"]
	// kiwi 값은 없는데 0이 초기화, lemon은 원래 있는 값 결과 구분이 안됨 -> 두 번째 ok 가 key가 존재 하는지 않아는지 true, false로 Return 함

	fmt.Println("ex1: ", value1)
	fmt.Println("ex1: ", value2, ok1)
	fmt.Println("ex1: ", value3, ok) //두 번째 리턴 값으로 키 존재 유무 확인

	if value, ok := map1["kiwi"]; ok {
		fmt.Println("ex2 :", value)
	} else {
		fmt.Println("ex2 :", "wiki is not exist!!")
	}

	if value, ok := map1["banana"]; ok {
		fmt.Println("ex2 :", value)
	} else {
		fmt.Println("ex2 :", "banana is not exist!!")
	}

	if _, ok := map1["kiwi"]; !ok {
		fmt.Println("ex2 :", "wiki is not exist!!")
	}

}
