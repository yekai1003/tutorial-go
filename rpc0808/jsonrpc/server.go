package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//服务结构体，空壳，载体
type Rect struct{}

type Params struct {
	Width, Height int
}

// 计算面积
func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height
	return nil
}

// 计算周长
func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func main() {
	rect := new(Rect)
	rpc.Register(rect) // 等价于 rpc.RegisterName("Rect",new(Rect))

	tcpaddr, err := net.ResolveTCPAddr("tcp4", "localhost:8080")
	if err != nil {
		log.Panic("failed to ResolveTCPAddr ", err)
	}
	//listen函数专门处理TCP
	listener, err := net.ListenTCP("tcp", tcpaddr)
	if err != nil {
		log.Panic("failed to Listen ", err)
	}
	fmt.Println("begin services...")

	for {
		//获取新的连接
		conn, err := listener.Accept()
		if err != nil {
			log.Panic("failed to Accept ", err)
		}
		//提供服务
		go jsonrpc.ServeConn(conn)
	}
}
