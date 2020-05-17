package main
import(
  "fmt"
  "section4/lib2"
)

var num int32

func init(){
  num = 30
  fmt.Println("main init start!!")
}

func main(){
  fmt.Println("10보다 큰 수 ?",lib2.CheckNum1(num))
}
