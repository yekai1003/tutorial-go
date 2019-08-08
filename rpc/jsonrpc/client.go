/*
   file    : client.go
   author  : yekai
   company : pdj(pdjedu.com)
*/
package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

//注意字段必须是导出
type Params struct {
	Width, Height int
}

func main() {
	//连接远程rpc服务
	//这里使用jsonrpc.Dial
	rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal(err)
	}
	ret := 0
	//调用远程方法
	//注意第三个参数是指针类型
	err = rpc.Call("Rect.Area", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
	err = rpc.Call("Rect.Perimeter", Params{50, 100}, &ret)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ret)
}
