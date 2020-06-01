package main

import (
	"fmt"
)

func sayhello(msg string) {
	defer func() {
		fmt.Println(msg)
	}()

	func() {
		fmt.Println("Hi!!")
	}()

}
func main() {
	sayhello("golang!!")
}
