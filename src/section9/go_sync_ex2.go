//동기화 고급1
package main

import (
	"fmt"
	"sync"
)

func main() {
	//고루틴 동기화 고급
	//대기 그룹
	//실행 흐흠 변경(고루틴이 종료 될 때 까지 대기 가능)
	//WaitGroup: Add(고루틴 추가), Done(작업 종료 알림), Wait(고루틴 종료 시 까지 대)
	//Add로 추가된 고루틴 개수와 Done으로 종료되는 알림 횟수 같아야 정확하게 동작한다.(Add == Done)

	wg := new(sync.WaitGroup) //Wait 그룹 생성

	for i := 0; i < 100; i++ {
		wg.Add(1) // for문이 한번 실행 될 때 1번 추가
		go func(n int) {
			fmt.Println("waitgroup:", n)
			wg.Done() // 추가된 1번이 끝남
		}(i)
	}

	//Add == done 횟수가 같아야 함
	wg.Wait()
	fmt.Println("WaitGroup End!!")

}
