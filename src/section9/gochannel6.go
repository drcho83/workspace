package main

import (
	"fmt"
)

func main() {
	//채널(channel)
	//close: 채널 닫기

	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- 7777
		}
	}()

	val1, ok1 := <-ch
	fmt.Println("Ex1 : ", val1, ok1)
	val2, ok2 := <-ch
	fmt.Println("Ex1 : ", val2, ok2)
	val3, ok3 := <-ch
	fmt.Println("Ex1 : ", val3, ok3)

	close(ch) // 채널 닫기
	val4, ok4 := <-ch
	fmt.Println("Ex1 : ", val4, ok4)

}
