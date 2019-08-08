package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	//1. 连接到服务器 rpc
	cli, err := rpc.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Panic("failed to Dial ", err)
	}
	defer cli.Close() //函数结束后关闭连接
	//2. 远程调用
	//func (client *Client) Call(serviceMethod string, args interface{}, reply interface{})
	var reply string
	//虽然执行了Call一个函数，但是实际已经通过网络发送了请求包，并且获取了响应包,gob序列化与反序列化
	if cli.Call("HelloService.Hello", "fuhongxue", &reply) != nil {
		log.Panic("failed to Call ", err)
	}

	fmt.Println(reply)
}
