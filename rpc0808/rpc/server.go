package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
)

//服务载体
type HelloService struct{}

/*
	1. 结构体可以导出
	2. 函数可以导出
	3. 函数只有2个参数，可以是go语言原生，也可以是自定义的
	4. 返回值是error
*/
//定义一个打招呼的服务
func (p *HelloService) Hello(user string, reply *string) error {
	*reply = "hello," + user
	return nil
}

/*
远程服务启动步骤：
 1. 定义结构体
 2. 定义结构体函数，服务函数满足前面的要求
 3. 注册服务
 4. 启动网络侦听
 5. 获得连接
 6. 为连接提供服务
*/

func main() {
	fmt.Println("hello world")
	//RegisterName(name string, rcvr interface{})
	//注册HelloService服务
	rpc.RegisterName("HelloService", new(HelloService))

	//开启侦听
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Panic("failed to Listen ", err)
	}

	//接收一个连接
	conn, err := listener.Accept()
	if err != nil {
		log.Panic("failed to Accept ", err)
	}

	rpc.ServeConn(conn)
}
