/*
   author:Yekai
   company:Pdj
   filename:10-channel3.go
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
	go func(out chan<- int) {
		for i := 0; i < 10; i++ {
			out <- i //向通道c1写入数据
			time.Sleep(time.Second * 1)
		}
		close(out)
	}(c1)
	//squarer
	go func(in <-chan int, out chan<- int) {
		for {
			num, ok := <-in //读c1数据
			if !ok {
				break
			}
			out <- num * num //将平方写入c2
		}
		close(out)

	}(c1, c2)
	//printer
	for {
		num, ok := <-c2
		if !ok {
			break
		}
		fmt.Println(num)
	}
}
