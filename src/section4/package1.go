package main

import (
	"fmt"
	"os"
)

func main() {
	//package는 코드 모듈화 및 재사용
	//응집도, 결합도
	//Go: 패키지 단위의 독립적이고 작은 단위로 개발 -> 작은 패키지를 결합해서 프로그래밍을 작성할 것을 권고
	//패키지 이름은 = 디렉토리 이름
	//같은 패키지 내 -> 소스파일들은 디렉토리명을 패키지 명으로 사용한다.
	//네이밍 규칙: 소문자 private, 대문자: public
	//go: main 패키지는 특별하게 인식 -> 컴파일러 공유 라이브러리가 아니라, 프로그램의 시작점 start point 로 인식

	var name string
	fmt.Println("이름은?: ")
	fmt.Scanf("%s", &name)                   //받는 부분
	fmt.Fprintf(os.Stdout, "Hi! %s\n", name) // command에서 해야할 때 Fprintf 사용

}
