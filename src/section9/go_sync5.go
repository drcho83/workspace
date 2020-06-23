package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//고루틴 동기화 객체
	//동기화 상태(조건) 메소드 사용
	//실행 흐름 제어
	//wait,notifiy, nofityALL : 기타 언어
	//Wait(실행 멈출 때), Signal(하나만 깨울 때), Broadcas (모두 깨울 때)

	//시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) // 비동기 버퍼 채널

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine Waiting: ", n)
			condition.Wait() //고루틴 대기(멈춤)
			fmt.Println("WARNING End", n)
			mutex.Unlock()
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-c
		//fmt.Println("received: ", <-c) // 두 번 채널을 통해서 받으므로 에러 발생
	}
	/*
		for i := 0; i < 5; i++ {
			mutex.Lock()
			fmt.Println("Wake Goroutine(Signal): ", i)
			condition.Signal() //힌개씩 깨움(모든 고루틴 생성 후)
			mutex.Unlock()
		}
	*/
	mutex.Lock()
	fmt.Println("wake goroutine(broadcast)")
	condition.Broadcast()
	mutex.Unlock()
	time.Sleep(2 * time.Second)

}
