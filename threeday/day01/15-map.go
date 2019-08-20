/*
   author:Yekai
   company:Pdj
   filename:15-map.go
*/
package main

import "fmt"

func main() {
	countryCapitalMap := make(map[string]string)

	// map插入key - value对,各个国家对应的首都
	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Roma"
	countryCapitalMap["China"] = "BeiJing"
	countryCapitalMap["India "] = "New Delhi"

	fmt.Println(countryCapitalMap["China"])
	//当key不存在时，直接打印不太优雅，可以使用下面的方法
	val, ok := countryCapitalMap["Japan"]
	if ok {
		fmt.Println("Japan's capital is", val)
	} else {
		fmt.Println("Japan's capital not in map")
	}

}
