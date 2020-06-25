//map 형태를 Json 변환
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//마샬링: 논리적 구조를 바이트로 변경
	data := make(map[string]interface{})
	data["name"] = "kim"
	data["age"] = 25

	//doc, _ := json.Marshal(data) // map을 Json 문서로 변환 // Marshal에서 리턴된 값은 byte 형식

	doc, _ := json.MarshalIndent(data, "", "  ")
	fmt.Println(doc)
	fmt.Println(string(doc))

}
