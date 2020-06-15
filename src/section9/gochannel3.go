package main

import (
	"fmt"
)

func main() {
	//채널(channel)
	//예제1(동기: 버퍼 미사용)

	ch := make(chan bool)
	cnt := 6

	go func() {
		for i := 0; i < cnt; i++ {
			ch <- true
			fmt.Println("Go: ", i)
			//time.Sleep(1 * time.Second) <- test~!!!!
		}
	}()

	for i := 0; i < cnt; i++ {
		<-ch
		fmt.Println("Main: ", i)
	}
}
