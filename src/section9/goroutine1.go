//고루팅(go routine) 기초(1)
package main

import (
	"fmt"
	"time"
)

func exe1() {
	fmt.Println("ex1 func start", time.Now())
	time.Sleep(1 * time.Second) //1초 microsecond, minute...
	fmt.Println("ex1 func end", time.Now())
}

func exe2() {
	fmt.Println("ex2 func start", time.Now())
	time.Sleep(1 * time.Second) //1초 microsecond, minute...
	fmt.Println("ex2 func end", time.Now())
}

func exe3() {
	fmt.Println("ex3 func start", time.Now())
	time.Sleep(1 * time.Second) //1초 microsecond, minute...
	fmt.Println("ex3 func end", time.Now())
}

func main() {
	//고루틴
	//타 언어의 스레드의 비슷한 기능
	//생성 방법 매우 간단, 리소스 매우 적게 사용
	//즉, 수많은 고루틴 동시 생성 실행 가능
	//비동기적 함수 루틴 실행(매우 적은 용량 차지) -> 채널을 통한 통신 가능
	//공유메모리 사용 시에 정확한 동기화 코딩 필요
	//싱글루민에 비해 항상 빠른 처리 결과는 아니다.

	//멀티 스레드의 장점과 단점
	//장점: 응답성 향상, 자원공유를 효율적으로 활용 사용, 작업이 분리되어 코드 간결
	//단점: 구현이 어려움, 테스트 및 디버깅 어령움, 전체프로세스의 사이드이펙스, 성능 저하, 동기화 코딩 반드시 숙지, 데드락
	//파이썬 스레드 추천!

	exe1() // 가장 먼저 실행(일반적인 실행 흐름)
	fmt.Println("main routine start", time.Now())
	go exe2() //나머지 exe2, exe3 은 실행되지 않는다. 그 이유는 데몬성 스레드로 main routine 이 끝나면 나머지 2,3 스레드가 실행 중 정지 된다.
	go exe3()
	time.Sleep(4 * time.Second)
	fmt.Println("main routine end", time.Now())
}
