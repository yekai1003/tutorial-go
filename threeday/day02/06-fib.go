/*
   author:Yekai
   company:Pdj
   filename:06-fib.go
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(time.Millisecond * 100) //启动一个打印goroutine
	fmt.Printf("\n%d\n", fib(45))
}

//计算斐波那契数列
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-2) + fib(x-1)
}

//此函数目的只是为了用户体验
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
