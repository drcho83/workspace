package main

import (
	"fmt"
)

func main() {
	//example1
Loop1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break Loop1
			}
			fmt.Println("ex1: ", i, j)
		}
	}

	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 2 && j == 4 {
				break //Loop1 을 설정하지 않았으므로, 이중 for 중 2단계 for문에서만 빠져 나오고 3부터 다시 실행된다.
			}
			fmt.Println("ex1: ", i, j)
		}
	}

	//example2
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue // 아래 fmt Print 구문이 실행 되지 않고, 조건 식으로 다시 돌아 감.
		}
		fmt.Println("ex2: ", i)
	}

Loop2: // Loop 아래는 반드시 for문이 와야 한다. Loop 레이블 밑에 관련없는 소스코드 시 에러
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 2 {
				continue Loop2
			}
			fmt.Println("ex3: ", i, j)
		}
	}
}
