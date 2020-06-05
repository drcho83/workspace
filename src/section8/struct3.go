//구조체 기본
package main

import (
	"fmt"
)

type Car struct {
	Color string
	name  string
}

func main() {
	c1 := Car{"red", "220d"}
	c2 := new(Car)
	c2.Color, c2.name = "white", "sonata"
	c3 := &Car{} // 빈구조체 선언
	c4 := &Car{"black", "520d"}

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(c3)
	fmt.Println(c4)
	fmt.Println()

}
