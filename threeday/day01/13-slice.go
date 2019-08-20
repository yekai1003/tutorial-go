/*
   author:Yekai
   company:Pdj
   filename:13-slice.go
*/
package main

import (
	"fmt"
)

func main() {
	a1 := [5]int{1, 2, 3, 4, 5} // a1 是一个数组
	s1 := a1[2:4]               //定义一个切片
	fmt.Println(a1)
	fmt.Println(s1)
	s1[1] = 100
	fmt.Println("after---------")
	fmt.Println(a1)
	fmt.Println(s1)
}
