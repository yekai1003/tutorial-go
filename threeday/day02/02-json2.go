/*
   author:Yekai
   company:Pdj
   filename:02-json2.go
*/
package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name     string   `json:"name"`
	Age      int      `json:"age"`
	Sex      string   `json:"sex"`
	GoodMan  bool     `json:"goodman,omitempty"`
	Farivate []string `json:"farivate"`
}

func main() {
	var p1 Person
	json_data := `{"name":"yekai","age":30,"sex":"man","goodman":false,"farivate":["drink","feidao"]}`
	json.Unmarshal([]byte(json_data), &p1)
	fmt.Println(p1)
}
