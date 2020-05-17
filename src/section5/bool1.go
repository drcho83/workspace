package main
import(
  "fmt"
)
func main(){
  //bool(boolean): 참고 거짓
  //조건부 논리 연산자랑 주로 사용: !(not), ||(or), &&(and)
  //암묵적 형변환 x: 0, Nil -> false 변환 없음 // 무조건 false or true 작성되어야 함

  var b1 bool =true
  var b2 bool =false
  b3 := true

  //b4 := 1 // error!! 암묵적 형변환 안됨. 에러 발생

  fmt.Println("ex1:", b1)
  fmt.Println("ex2:", b2)
  fmt.Println("ex3:", b3)

  fmt.Println("")
  fmt.Println("ex4:", b3 == b3)
  fmt.Println("ex5:", b1 && b3)
  fmt.Println("ex6:", b1 || b2)
  fmt.Println("ex7:", !b3)

  /*
  if b4 {
    fmt.Println("ex8: true!!") // error!! 암묵적 형변환 안됨. 에러 발생
  }
  */

}
