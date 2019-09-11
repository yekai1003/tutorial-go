/*
   author:Yekai
   company:Pdj
   filename:14-open.go
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fd, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("failed to openfile ", err)
		return
	}
	defer fd.Close() //延迟关闭
	fd.WriteString("hello,yekai\n")
	fd.Seek()
}
