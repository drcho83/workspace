package main
import(
  "fmt"
  checkUp "section4/lib"
  _ "section4/lib2" // 사용하지 않지만 빈 식별자를 사용해서 일단 남겨 놓을 때 _ 사용
)

func main(){
  //패키지 접근 제어
  //별칭 사용
  //빈 식별자 사용
  fmt.Println("10보다큰수?: ",checkUp.CheckNum(11))
}
