/*
   author:Yekai
   company:Pdj
   filename:15-readdir.go
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	fd, err := os.Open("./") //打开当前目录
	if err != nil {
		fmt.Println("failed to openfile ", err)
		return
	}
	//参数<0 代表要读取全部，大于0也是代表要读取部分
	infos, err := fd.Readdir(-1)
	if err != nil {
		fmt.Println("failed to Readdir ", err)
		return
	}
	//遍历读取目录的结果
	for _, v := range infos {
		fmt.Printf("name = %s, isdir = %t, size = %d\n", v.Name(), v.IsDir(), v.Size())
	}
}
