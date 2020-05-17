//go 초기화 함수1
package main
import(
  "fmt"
)

func init(){
  //main 함수보다 먼저 실행 되며 한번 실행된다.
  //한번만 되면 되므로 프로그램 상태 체크 1회용?
  //패키지 로드 시에 가장 먼저 호출
  //가장 먼저 초기화 되는 작업 작성 시 유용하다.
  fmt.Println("init Method Start")
}

func main(){
  fmt.Println("Main Method Start!!")
}
