package main

import (
	"fmt"
	"log"
	"math"
	"time"
)

//예외 처리 구조체

type PowError struct {
	time    time.Time // 에러 발생 시간
	value   float64   // 파라미터
	message string    // 에러 메시지
}

func (e PowError) Error() string { // 리시버를 통해서 Error 매소드 구현 Return 타입은 string // Error() <<- 인터페이스 사용
	return fmt.Sprintf("[%v]Error - Input Value(value: %g) - $s", e.time, e.value, e.message)
}

func Power(f, i float64) (float64, error) {
	if f == 0 {
		return 0, PowError{time: time.Now(), value: f, message: "0은 사용할 수 없습니다."}
	}
	if math.IsNaN(f) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다."}
	}
	if math.IsNaN(i) {
		return 0, PowError{time: time.Now(), value: f, message: "숫자가 아닙니다."}
	}
	return math.Pow(f, i), nil
}

func main() {
	//example1
	v, err := Power(10, 3)
	if err != nil {
		log.Fatal(err)
		//log.Fatal(err.Error())
	}
	fmt.Println("ex1:", v)

	t, err := Power(0, 3)
	if err != nil {
		//log.Fatal(err)
		log.Fatal(err.Error())
		//fmt.Println(err.(PowError).value)
		//fmt.Println(err.(PowError).message)
		//fmt.Println(err.(PowError).time)
	}
	fmt.Println("ex2:", t)

}
