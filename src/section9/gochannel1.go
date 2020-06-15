//채널(channel) 기초 (1) // 동기식 채널 활용 법
package main

import (
	"fmt"
	"time"
)

func work1(v chan int) { // 채널인데 int형 채널
	fmt.Println("work1 : S -->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work1 : E -->", time.Now())

	v <- 1 // 1을 채널로 송신
}

func work2(v chan int) { // 채널인데 int형 채널
	fmt.Println("work2 : S -->", time.Now())
	time.Sleep(1 * time.Second)
	fmt.Println("work2 : E -->", time.Now())

	v <- 2 // 2을 채널로 송신
}

func main() {
	//채널
	//고루틴간의 상호 정보(데이터) 교환 및 실행 흐룸 동기화 위해 사용 <<- 동기식 채널 생성 (동기식, 레퍼런스!!!!!!!!)
	//실행 흐름 제어 가능 (동기, 비동기) -> 채널 자체를 일반 변수로 선언 후 사용 가능
	//데이터 전달 자료형 선언 후 사용 (지정 된 타입만 주고받을 수 있음)
	//그러나, interface{}전달을 통해서 자료형 상관없이 전송 및 수신가능
	//값을 전달 (복사 후: bool, int등), 포인터(슬라이스, 맵) 등을 전달시 주 <- 첨조형으로 원본값 수정될 수 있음 (동기화 사용 = mutex)
	//멀티프로세싱 처리에서 교착상태(경쟁) 주의!!!
	//<-,-> (채널 <- 데이터) 데이터를 채널로 송신 의미, (변수 -< 채널) 채널을 변수로 보내는 것을 수신이라고 표현
	//example1
	fmt.Println("Main : S -->", time.Now())
	//var c chan int
	// c = make(chan int) 정석적인 방법
	v := make(chan int) // int형 채널 선언
	go work1(v)
	go work2(v)
	<-v
	<-v

	//time.sleep 으로 Delay를 강제로 줄 필요가 없다. 동기화 식이기 때문에 수신이 완료될 때까지 대기한다.

	fmt.Println("Main : E -->", time.Now())
}
