/*
   author:Yekai
   company:Pdj
   filename:05-goroutine.go
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("begin call goroutine")
	//启动goroutine
	go func() {
		fmt.Println("I am a goroutine!")
	}()
	fmt.Println("end call goroutine")
	time.Sleep(time.Second * 1) //睡1s
}
