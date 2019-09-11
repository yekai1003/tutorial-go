/*
   author:Yekai
   company:Pdj
   filename:01-json.go
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

// type Person struct {
// 	Name     string
// 	age      int
// 	Sex      string
// 	GoodMan  bool
// 	Farivate []string
// }

func main() {
	p1 := Person{"yekai", 30, "man", false, []string{"drink", "feidao"}}
	data, err := json.Marshal(p1)
	if err != nil {
		fmt.Println("failed to Marshal data ", err)
		return
	}
	fmt.Println(string(data)) //string(data)是类型强制转会
}
