//라이브러리 접근제어(1)
package lib2

import(
  "fmt"
)

func CheckNum1(c int32) bool{ // 소문자로 변경할 경우 함수가 없다고 인식하므로 대문자로 정의 되어야 한다.
  return c >100
}

func CheckNum2(c int32) bool{ // 소문자로 변경할 경우 함수가 없다고 인식하므로 대문자로 정의 되어야 한다.
  return c >1000
}

func init(){
    fmt.Println("package init Start")
}
