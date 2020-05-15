//상수1
package main

import "fmt"

func main() {
	//상수를 선언하면 이 값은 변화되지 않는다. 불변
	//const 사용 초기화, 한 번 선언 후 값 변경 불가, 고정 된 값으로 관리 용
	const a string = "test1" // 초기화 바로 진행되어야 함
	const b = "test2"
	const c int32 = 10 * 10 // 100할당
	//const d =getHeigh() // Comple Error 발생, getHeigh가 동일한 값을 가져 오는지 확답할 수 없기 때문에 에러가 남
	const e = 356.6
	const f = false
	/*
	  에러 발생
	  const g string
	  g = "test3" <<-- 상수는 선언과 동시에 선언되어야 하므로 에럴 발생됨
	*/

	fmt.Println("a: ", a, "b: ", b, "c: ", c, "e: ", e, "f: ", f)
}
