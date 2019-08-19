/*
   author:Yekai
   company:Pdj
   filename:08-func.go
*/

package main

import (
	"fmt"
)

func main() {
	//函数调用，同时获得2个返回值
	sum, sub := add_sub(32, 21)
	fmt.Println(sum, sub)
	//获得函数指针，此时addsubptr相当于 func add_sub(a int, b int) (int, int)
	addsubptr := add_sub
	//通过函数指针的方式调用
	sum1, sub1 := addsubptr(1, 2)
	fmt.Println(sum1, sub1)
}

func add_sub(a int, b int) (int, int) {
	return a + b, a - b
}
