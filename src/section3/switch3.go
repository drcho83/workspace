package main

import (
	"fmt"
)

func main() {
	//example1
	a := 30 / 15
	switch a {
	case 2, 4, 6:
		fmt.Println("a ->", a, "는 짝수")
	case 1, 3, 5:
		if a > 1 {
			fmt.Println("a ->", a, "1보다 큰 홀수")
		}
		fmt.Println("a ->", a, "는 홀수")
	}

	//example2
	switch e := "go"; e {
	case "Java":
		fmt.Println("Java")
		fallthrough
	case "go":
		fmt.Println("go")
		fallthrough // 조건에 맞는 문장 아래에 반드시 실행되어야 하는 로직이 있을 경우 fallthrough를 넣어야지 반드시 실행이 됨
	case "python":
		fmt.Println("python")
		fallthrough
	case "rubby":
		fmt.Println("rubby")
		fallthrough
	case "c":
		fmt.Println("c")
	}
}
