//고루팅(go routine) 기초(1)
package main

import (
	"fmt"
	"time"
)

func exe(name string) {
	fmt.Println(name, "stat :", time.Now())
	for i := 1; i < 1000; i++ {
		fmt.Println(name, ">>>>>", i)
	}
	fmt.Println(name, "end :", time.Now())
}
func main() {
	//고루틴
	exe("t1")
	fmt.Println("main routine start", time.Now())
	go exe("t2")
	go exe("t3")
	time.Sleep(4 * time.Second)
	fmt.Println("main routine end", time.Now())
}
