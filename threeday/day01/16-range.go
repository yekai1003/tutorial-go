/*
   author:Yekai
   company:Pdj
   filename:16-range.go
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

	//遍历map
	for k, v := range countryCapitalMap {
		fmt.Println(k, "'s capital is", v) //k,v分别是map的key和val
	}
	//遍历数组，如果不想获得
	a := []int{10, 20, 30, 40, 50}
	for k, v := range a {
		fmt.Printf("a[%d]=%d\n", k, v) //k代表数组下标，v代表该元素值
	}

}
