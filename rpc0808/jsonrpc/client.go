package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

//结构体定义应与服务端一致
type Params struct {
	Width, Height int
}

func main() {
	//使用jsonrpc进行Dial连接
	cli, err := jsonrpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Panic("failed to Dial", err)
	}

	defer cli.Close()
	ret := 0
	if cli.Call("Rect.Area", Params{50, 120}, &ret) != nil {
		log.Panic("failed to Call Rect.Area ", err)
	}
	fmt.Println("call Rect.Area ===", ret)
	if cli.Call("Rect.Perimeter", Params{50, 120}, &ret) != nil {
		log.Panic("failed to Call Rect.Area ", err)
	}
	fmt.Println("call Rect.Perimeter ===", ret)
}
