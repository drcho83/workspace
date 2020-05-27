package main

import (
	"fmt"
)

func main() {
	//맵 조회 및 순회 (iterator)
	//example1
	map1 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.net",
		"google": "http://google.net",
	}
	fmt.Println("ex1 :", map1["google"])
	fmt.Println("ex1 :", map1["daum"])
	fmt.Println()
	//example2 순위
	for k, j := range map1 {
		fmt.Println("ex2 :", k, j)
	}

	fmt.Println()

	for _, v := range map1 {
		fmt.Println("ex2: ", v)
	}



}
