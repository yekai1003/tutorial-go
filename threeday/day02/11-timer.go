/*
   author:Yekai
   company:Pdj
   filename:11-timer.go
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.NewTimer(time.Second * 1)
	data := <-timer.C
	fmt.Println(data)
	timer.Stop() //停止定时器
}
