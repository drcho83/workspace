package main

import "fmt"

func main() {
	//맵: 해시테이블, 딕셔너리(파이썬), key-vaule 자료 저장
	//레퍼런스 타입(참조값 전달)
	//비교 연산자 사용 불가능(참조 타입이므로)
	//특징: 참조타입(key)로 사용 불가능, 값(value)으로 모든 타입 사용 가능
	//make 함수 및 축약(리터럴)로 초기화 가능
	//순서 없음

	//example1
	var map1 map[string]int = make(map[string]int) // 정석 key , value 형을 각각 선언해야함
	var map2 = make(map[string]int)
	map3 := make(map[string]int)

	fmt.Println("ex1 :", map1)
	fmt.Println("ex1 :", map2)
	fmt.Println("ex1 :", map3)
	fmt.Println()

	//example2
	map4 := map[string]int{} //json 형태와 비슷 a = {"a":1,"b":n}
	map4["apple"] = 25
	map4["banana"] = 40
	map4["orange"] = 33

	map5 := map[string]int{
		"apple":  15,
		"banana": 15,
		"orage":  15,
	}

	map6 := make(map[string]int, 10)
	map6["apple"] = 25
	map6["banana"] = 40
	map6["orange"] = 33

	fmt.Println("ex2 : ", map4)
	fmt.Println("ex2 : ", map5)
	fmt.Println("ex2 : ", map6)
	fmt.Println("ex2 : ", map6["orange"])
	fmt.Println("ex2 : ", map6["apple"])

}
