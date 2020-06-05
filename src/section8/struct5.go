//익명 구조체
package main

import (
	"fmt"
)

func main() {
	//구조체 익명 선언 및 활용
	//example1

	car1 := struct{ name, color string }{"520d", "red"} //type aaa struct {aaa int, ~~~~}

	fmt.Println("ex1 :", car1)
	fmt.Printf("ex1 : %#v\n", car1)

}
