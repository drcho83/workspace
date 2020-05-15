//변수3 - 집중 Go 특징
package main

import "fmt"

func main() {
	//짧은 선언
	// 반드시 안에서만 사용 (전영 사용 안됨), 선언 후 재할당 하면 예외 발생
	// 안에서만 일회성으로만 사용하는 용도 (명시적 활용)
	// 주로 제한된 범위의 함수내에서 사용할 경우 코드 가독성을 높일 수 있음 ==> 추천
	shortVar1 := 3
	shortVar2 := "Test"
	shortVar3 := false

	//shortVar1 :=3  <<- 에러 발생

	fmt.Println("shortval1:", shortVar1, "shortvar2: ", shortVar2, "shorvar3: ", shortVar3)

	//example
	if i := 10; i < 11 {
		fmt.Print("Short Variable Test Success!!")
	}

}
