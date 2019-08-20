/*
   author:Yekai
   company:Pdj
   filename:14-slice2.go
*/
package main

import (
	"fmt"
)

func main() {
	var s1 []int       //定义切片s1
	s1 = append(s1, 1) //追加，注意s1必须接收追加结果
	s1 = append(s1, 2)
	s1 = append(s1, 3, 4, 5) //可以一次追加多个
	printSlice(s1)
	s2 := make([]int, 3)
	printSlice(s2)
	s2 = append(s2, 4) //当超过容量的时候，容量会以len*2的方式自动扩大
	printSlice(s2)
}

func printSlice(s []int) {
	fmt.Printf("len = %d, cap = %d, s = %v\n", len(s), cap(s), s)
}
