/*
   author:Yekai
   company:Pdj
   filename:07-for.go
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	//计算1+2+3+……+100 = ？
	//第一种方式
	sum := 0
	i := 0
	for i = 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println("sum is ", sum)
	//第二种方式
	i = 1
	sum = 0
	for i <= 100 {
		sum += i
		i++
	}
	fmt.Println("sum is ", sum)
	//死循环 - 开启刷屏模式
	for {
		fmt.Println("heihei")
		time.Sleep(time.Second * 1) //每次执行睡眠1s
	}
}
