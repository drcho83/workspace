package main

import (
	"fmt"
)

func main() {
	//맵 값 변경 및 삭제
	map3 := map[string]string{
		"daum":   "http://daum.net",
		"naver":  "http://naver.net",
		"google": "http://google.net",
		"home1":  "http://test1.com",
	}
	fmt.Println("ex1: ", map3)

	map3["home2"] = "http://test2.com" //자동 추가 됨
	fmt.Println("ex1:", map3)

	map3["home2"] = "http://test2-2.com" // 키가 존재할 경우 자동 수정됨
	fmt.Println("ex1:", map3)
	fmt.Println()

	//example2 (삭제)

	delete(map3, "home2")
	fmt.Println("ex1:", map3)

  delete(map3, "google")
  fmt.Println("ex1:", map3)



}
