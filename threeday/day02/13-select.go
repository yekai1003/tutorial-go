/*
   author:Yekai
   company:Pdj
   filename:13-select.go
*/
package main

import (
	"fmt"
	"os"
	"time"
)

func launch() {
	fmt.Println("发射!")
}

func main() {
	ticker := time.NewTicker(time.Second)
	fmt.Println("开始倒计时准备发射，按回键可以取消发射！")
	num := 5
	chan_stdin := make(chan string)
	go func(out chan<- string) {
		data := make([]byte, 10)
		os.Stdin.Read(data) //该goroutine读也会阻塞
		out <- "cancel"
	}(chan_stdin)
	for {
		select {
		case <-ticker.C:
			fmt.Println(num)
			num--
		case <-chan_stdin:
			fmt.Println("取消发射！")
			return
		}
		if num == 0 {
			ticker.Stop()
			break
		}
	}

	launch() //发射火箭
}
