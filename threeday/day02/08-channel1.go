/*
   author:Yekai
   company:Pdj
   filename:08-channel1.go
*/
package main

import (
	"fmt"
	"time"
)

var c chan string

func reader() {
	msg := <-c
	fmt.Println("I am reader,", msg)
}

func main() {
	c = make(chan string)
	go reader()
	fmt.Println("begin sleep")
	time.Sleep(time.Second * 3) //睡眠3s为了看执行效果
	c <- "hello"
	time.Sleep(time.Second * 1) //睡眠1s为了看执行效果
}
