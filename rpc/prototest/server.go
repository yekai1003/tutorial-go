package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"tutorial-go/rpc/proto"
)

type HelloService struct{}

func (p *HelloService) Hello(request *goprotoc.MsgValue, reply *goprotoc.MsgValue) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func (p *HelloService) Method(request *goprotoc.MethodVal, reply *goprotoc.MethodVal) error {
	//reply.Value = "hello:" + request.GetValue()
	switch request.GetMethod() {
	case "add":
		reply.Reply = request.GetX() + request.GetY()
	case "mul":
		reply.Reply = request.GetX() * request.GetY()
	case "div":
		if request.GetY() == 0 {
			return errors.New("params err")
		}
		reply.Reply = request.GetX() / request.GetY()
	case "sub":
		reply.Reply = request.GetX() - request.GetY()
	default:
		return errors.New("func err")
	}
	return nil
}

func main() {
	rpc.RegisterName("HelloService", new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	conn, err := listener.Accept()
	if err != nil {
		log.Fatal("Accept error:", err)
	}

	rpc.ServeConn(conn)
}
