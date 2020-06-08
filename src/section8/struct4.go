//구조체 기본(4)
package main

import "fmt"

func main() {
	//구조체 익명 선언 및 활용

	//예제1
	car1 := struct{ name, color string }{"520d", "red"}

	fmt.Println("ex1 : ", car1)
	fmt.Printf("ex1 : %#v\n", car1)

	//예제2
	cars := []struct{ name, color string }{{"520d", "red"}, {"220d", "white"}, {"420d", "black"}}
	for _, c := range cars {
		fmt.Printf("(%s, %s) ----- (%#v)\n", c.name, c.color, c)
	}

}
