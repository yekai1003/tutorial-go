package main

import (
	"log"
	"net"
	"net/rpc"
)

type HelloService struct{}

func (p *HelloService) Hello(user string, reply *string) error {
	*reply = "hello:" + user
	return nil
}

func main() {

	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
