//사용자 패키지 작성 및 문서화 예제

package main

import (
  "fmt"
  oper "section12/arithmetic" // alias 사용 (패키지 중복 또는 약사로 사용)
)

func main(){
  //사용자 패키지 작성 & GO 문서화
  //기준은 GOPATH/src
  //패키지 폴더명 (디렉토리명) 명확하게 지정
  //패키지 파일의 package 이름 으로 사용한다. --> 길면 alias 를 사용
  //package main 을 제외하고 package 문서에 등록
  //지금까지 우리는 패키지를 사용해 왔다.
  //기본적으로 GOROOT의 패키지 폴더를 검색하고 없으면 GOPATH의 패키지 검색(src/pkg)
  //go install 명령어 실행의 경우 -> gopath/pkg 에 등록 (문서화)
  //godoc -http=:6060(임의의포트) -> pkg 이동 -> 본인 패키지 메소드 및 주석 확인 (패키지, 타입, 메소드) 주석


  //패키지 사용 예제(사칙연산)
  nums := oper.Numbers{100,10}
  fmt.Println("package Used(1): ", nums.Plus())
  fmt.Println("package Used(2): ", nums.Minus())
  fmt.Println("package Used(3): ", nums.Multi())
  fmt.Println("package Used(4): ", nums.Divide())
  fmt.Println("package Used(5): ", nums.SquarePlus())
  fmt.Println("package Used(6): ", nums.SquareMinus())

}
