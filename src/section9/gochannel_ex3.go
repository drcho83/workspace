//채널 심화 1
package main

import (
	"fmt"
)

func receiveOnly(cnt int) <-chan int {
	sum := 0
	tot := make(chan int)
	go func() {
		for i := 1; i <= cnt; i++ {
			sum += i
		}
		tot <- sum
		tot <- 777
		tot <- 777
		close(tot)
	}()
	return tot
}

func total(c <-chan int) <-chan int {
	tot := make(chan int)
	go func() {
		a := 0
		for i := range c {
			a = a + i
		}
		tot <- a
	}()
	return tot
}

func main() {
	//채널 또한 함수의 반환 값으로 사용 가능
	//example1
	c := receiveOnly(100) // 채널 반환
	output := total(c)
	//output <- 777 // error 발생 output은 수신 전용 변수이므로 값을 보낼 시 에러 발생

	fmt.Println("ex1 :", <-output)
}
