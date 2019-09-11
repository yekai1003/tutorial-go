/*
   author:Yekai
   company:Pdj
   filename:09-channel2.go
*/
package main

import (
	"fmt"
	"time"
)

var c1 chan int
var c2 chan int

func main() {
	c1 = make(chan int)
	c2 = make(chan int)
	//counter
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i //向通道c1写入数据
			time.Sleep(time.Second * 1)
		}
		close(c1)
	}()
	//squarer
	go func() {
		for {
			num, ok := <-c1 //读c1数据
			if !ok {
				break
			}
			c2 <- num * num //将平方写入c2
		}
		close(c2)

	}()
	//printer
	for {
		num, ok := <-c2
		if !ok {
			break
		}
		fmt.Println(num)
	}
}
