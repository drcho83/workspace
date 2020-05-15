package main

import "fmt"

func main() {
	const (
		_ = iota
		A
		_
		C
	)
	fmt.Println(A, C)

	const (
		_ = iota + 0.75*2
		Default
		Silver
		gold
		Platium
	)

	fmt.Println("D:", Default)
	fmt.Println("S:", Silver)
	fmt.Println("G:", gold)
	fmt.Println("P:", Platium)
}
