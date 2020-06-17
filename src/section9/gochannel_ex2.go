//채널 심화 1
package main

import (
	"fmt"
)

func sum(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 0; i < cnt; i++ {
			sum += i
		}
		tot <- sum
	}()
	return tot
}

func main() {
	//채널 또한 함수의 반환 값으로 사용 가능
	//example1
	c := sum(100)
	fmt.Println("ex1: ", <-c)
}
