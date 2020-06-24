//동기화 고급1
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	//고루틴 동기화 고급
	//원자성 사용 -> all or nothing -> 기능적으로 분할 불가능한 완전 보증된 일련의 조작, 모두 성공하거나, 모두 실패
	//모든 조작이 완료 될 때까지 다른 프로세스 개입 불가
	//sync/atomic에서 원자적 연산자 제공
	//https://golang.org/pkg/sync/atomic 에서 계열 확인 가능
	//주로 공용 변수에 관한 계산 사용
	//lock / unlock으로 해도 됨

	// 원자성 사용 사용 안할 경우 예제
	runtime.GOMAXPROCS(runtime.NumCPU())
	var cnt int64 = 0

	wg := new(sync.WaitGroup) //Wait 그룹 생성

	for i := 0; i < 5000; i++ {
		wg.Add(1) // for문이 한번 실행 될 때 1번 추가
		go func(n int) {
			cnt += 1
			wg.Done() // 추가된 1번이 끝남
		}(i)
	}

	for i := 0; i < 2000; i++ {
		wg.Add(1) // for문이 한번 실행 될 때 1번 추가
		go func(n int) {
			cnt -= 1
			wg.Done() // 추가된 1번이 끝남
		}(i)
	}
	//Add(7000) == done(7000) 횟수가 같아야 함
	wg.Wait()
	fmt.Println("WaitGroup End!! cnt?? >>> ", cnt)

}
