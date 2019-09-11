/*
   author:Yekai
   company:Pdj
   filename:03-json-map.go
*/
package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	json_map := make(map[string]interface{}) //构造一个map，注意map的值类型应为接口类型
	json_data := `{"name":"yekai","age":30,"sex":"man","goodman":false,"farivate":["drink","feidao"]}`
	err := json.Unmarshal([]byte(json_data), &json_map)
	if err != nil {
		fmt.Println("failed to Unmarshal", err)
		return
	}
	fmt.Println(json_map)
	//下面用类型断言
	name, ok := json_map["name"].(string)    //断言name是string类型，ok是指示器
	fmt.Println(ok, name)                    //正常应该做一下ok为真的判断
	isgood, ok := json_map["goodman"].(bool) //bool型断言
	fmt.Println(ok, isgood)
	farivates, ok := json_map["farivate"].([]interface{}) //对于切片的类型来说，先断言为泛型切片
	fmt.Println(ok, farivates[0])                         //当然也可以继续对farivates[0]进行断言
}
