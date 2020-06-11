package main

import (
	"fmt"
	"runtime" // 운영체제의 정보를 접근 할 수 있는 pkg
	"time"
)

func main() {
	//고루팅(goroutine)
	//클로저 사용 예제
	//example1
	runtime.GOMAXPROCS(2)
	s := "Goroutine Clousure :" // 캡쳐

	for i := 0; i < 1000; i++ {
		go func(n int) {
			fmt.Println(s, n, " - ", time.Now())
		}(i) //반복문 클로저는 일반적으로 즉시 실행 but 고 루 글로저는 가장 나중에 실행 (반복문 종료 후 실행)
	}
	time.Sleep(5 * time.Second)
}
