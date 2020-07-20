package main

import (
	"fmt"
	"os"
)

func fileOpen(filename string) {
	//defer 함수
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("File Open Error:", r)
		}
	}()

	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("ex1: ", f.Name())
	}

	defer f.Close()
}

func main() {
	fileOpen("undefinded.txt")
	fmt.Println("End Main")
}
