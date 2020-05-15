package main

import "fmt"

func main() {
	//swich 뒤 표현식 생략 가능
	//case 뒤 표현식 사용 가능
	//자동 break 때문에 fallthrouth whswo
	//type 분기 -> 값이 아닌 변수 type으로 분기 가능

	//example1
	a := -7
	switch {
	case a < 0:
		fmt.Println(a, "는 음수")
	case a == 0:
		fmt.Println(a, "0")
	case a > 0:
		fmt.Println(a, "는 양수")
	}

	//example2
	switch b := 27; {
	case b < 0:
		fmt.Println(b, "는 음수")
	case b == 0:
		fmt.Println(b, "0")
	case b > 0:
		fmt.Println(b, "는 양수")
	}

	//example3
	switch c := "go"; c { //잛은 선언 후  c를 선언하면 같다로 매칭이 된다!!!
	case "go": //case c=="go"
		fmt.Println("GO!~!")
	case "java":
		fmt.Println("JAVA!~!")
	default:
		fmt.Println("일치 하는 값 없음")
	}

	switch c := "go"; c + "lang" {
	case "golang":
		fmt.Println("golang")
	case "Java":
		fmt.Println("Java")
	default:
		fmt.Println("nothing")
	}

	//example5
	switch i, j := 20, 30; {
	case i < j:
		fmt.Println("i는 j보다 작다")
	case i == j:
		fmt.Println("i와 j는 같다")
	case i > j:
		fmt.Println("i는 j보다 크다")
	}

}
