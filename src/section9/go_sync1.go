//고루틴 동기화 기초
package main

import (
	"fmt"
	"runtime"
)

//구조체 선언(공유 데이터)
type count struct {
	num int
}

func (c *count) increament() {
	c.num += 1
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	//고루틴 동기화 예정
	//실행 흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장 중요

	//동기화 사용하지 않은 경우 example1
	//시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increament()
			done <- true
			runtime.Gosched() // CPU 양보
		}()
	}

	for i := 1; i <= 10000; i++ {
		<-done
	}

	c.result()

}
