/*
   author:Yekai
   company:Pdj
   filename:11-array.go
*/

package main

import (
	"fmt"
)

func main() {
	var a1 [5]int = [5]int{1, 2, 3, 4}
	fmt.Println(a1)
	a1[4] = 6
	fmt.Println(a1)
	s1 := [4]string{"yekai", "fuhongxue", "luxiaojia"} //元素个数不能超过数组个数
	fmt.Println(s1)

}
