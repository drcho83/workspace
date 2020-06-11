//고루팅(go routine) // CPU 개수 지정
package main

import (
	"fmt"
	"math/rand"
	"runtime" // 운영체제의 정보를 접근 할 수 있는 pkg
	"time"
)

func exe(name int) {
	r := rand.Intn(100) // 100미만의 랜럼하게 수를 가져 온다.
	fmt.Println(name, "start :", time.Now())
	for i := 0; i < 100; i++ {
		fmt.Println(name, ">>>>>>>", r, i)
	}
	fmt.Println(name, "end :", time.Now())
}

func main() {
	//고루틴
	//멀티코어 (다중 CPU) 최대한 활용
	runtime.GOMAXPROCS(runtime.NumCPU())                       // 현 시스템의 CPU 코어 개수 반환 후 설정
	fmt.Println("Current System CPU: ", runtime.GOMAXPROCS(0)) // 설정 값 출력

	//example1
	fmt.Println("Main Routine Start: ", time.Now())
	for i := 0; i < 100; i++ {
		go exe(i) // 고루틴 100개 생성
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Main Routine End: ", time.Now())
}
