/*
   file    : server.go
   author  : yekai
   company : pdj(pdjedu.com)
*/
package main

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//注意字段必须是导出
type Params struct {
	Width, Height int
}

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	rect := new(Rect)
	//注册rpc服务
	rpc.Register(rect)
	//获取tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080")
	if err != nil {
		log.Panic("failed to ResolveTCPAddr ", err)
	}
	//监听端口
	tcplisten, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		log.Panic("failed to ListenTCP ", err)
	}
	for {
		conn, err := tcplisten.Accept()
		if err != nil {
			continue
		}
		//使用goroutine单独处理rpc连接请求
		//这里使用jsonrpc进行处理
		go jsonrpc.ServeConn(conn)
	}
}
