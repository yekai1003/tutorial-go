/*
   author:Yekai
   company:Pdj
   filename:04-panic.go
*/
package main

import (
	"fmt"
)

func panic_recover() {
	//recover()
	defer func() { //延迟定义捕获
		fmt.Println("defer recover panic") //会打印
		recover()
	}()
	panic("game over") //抛出错误

	fmt.Println("panic_recover run over") //不会打印
}

func main() {
	//recover()
	panic_recover()
	fmt.Println("heihei,game not over!") //捕获成功则会打印本句
}
