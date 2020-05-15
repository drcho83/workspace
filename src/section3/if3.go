package main

import "fmt"

func main() {
	i := 100
	if i >= 120 {
		fmt.Println("120이상")
	} else if i >= 100 && i < 120 {
		fmt.Println("100이상 120미만")
	} else if i < 100 && i >= 50 {
		fmt.Println("50이상 100미만")
	} else {
		fmt.Println("50 미만")
	}
}
