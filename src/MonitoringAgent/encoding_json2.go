//Json 형태로 GO 받기
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	//마샬링: 논리적 구조를 바이트로 변경

	doc := `
	{
		"name": "maria",
		"age": 10
	}
	`
	data := make(map[string]interface{})
	json.Unmarshal([]byte(doc), &data)
	fmt.Println(data["name"], data["age"])

}
