package main
import(
  "fmt"
  "section4/lib2"
)
func init(){
    fmt.Println("init Method Start")
}
func init(){
    fmt.Println("init2 Method Start")
}
func init(){
    fmt.Println("init3 Method Start")
}
//한 파일에 여러개 초기화 함수가 있어서 실행 됨. 순서대로 실행됨
//main이 없으면 실해 오류 발생
//package 호출을 위한 go 파일 생성 시 main 함수 없이 함수로만 구성 되어 있음
//package 호출을 위한 파일에 init 함수를 넣을 경우 해당 package import 시 init 이 먼저 실행된다.

func main(){
    fmt.Println("main Method Start")
    fmt.Println(lib2.CheckNum1(100))
}
