//채널 심화 1
package main

import (
	"fmt"
	"time"
)

func sendOnly(c chan<- int, cnt int) {
	for i := 0; i < cnt; i++ {
		c <- i
	}
	c <- 777
	//fmt.Println(<-c) Error 발생.. 발신 전용 함수
}

func receiveOnly(c <-chan int) {
	for i := range c {
		fmt.Println("recived: ", i)
	}
	fmt.Println(<-c)
}

func main() {
	//채널
	//함수 등의 매개변수에 수신 및 발신 전용 채널 지정 가능
	//전용 채널 설정 후 방향이 다를 경우 예외 발생
	//발신 전용 channel <- 데이터형
	//수신 전용 <- channel
	//매개 변수를 통해서 번용 채널 확인할 수 있다.
	//채널 또한 함수의 반환 값으로 사용 가능
	//example
	c := make(chan int)
	go sendOnly(c, 10) //발신 전용
	go receiveOnly(c)  //수신 전용
	time.Sleep(2 * time.Second)
}
