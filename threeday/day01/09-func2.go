/*
   author:Yekai
   company:Pdj
   filename:09-func2.go
*/

package main

import (
	"fmt"
)

func main() {
	a := math(10, 20, add) //传入add函数，求和
	b := math(10, 20, sub) //传入sub函数，求差
	fmt.Println(a, b)
}

func add(a int, b int) int {
	return a + b
}
func sub(a int, b int) int {
	return a - b
}

//函数作为特殊的类型也可以当作参数,调用时要求f参数必须是 func(a, b int) int 这样类型的函数
func math(a, b int, f func(a, b int) int) int {
	return f(a, b)
}
