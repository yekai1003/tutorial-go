/*
   file    : server.go
   author  : yekai
   company : pdj(pdjedu.com)
*/
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

func (p *HelloService) Method(request *goprotoc.MethodVal, reply *goprotoc.MethodReply) error {
	//reply.Value = "hello:" + request.GetValue()
	reply.Code = 0
	reply.Msg = "OK"
	switch request.GetMethod() {
	case "add":
		reply.Reply = request.GetX() + request.GetY()
	case "mul":
		reply.Reply = request.GetX() * request.GetY()
	case "div":
		if request.GetY() == 0 {
			reply.Code = 100
			reply.Msg = "Divisor is zero"
			return errors.New("params err")
		}
		reply.Reply = request.GetX() / request.GetY()
	case "sub":
		reply.Reply = request.GetX() - request.GetY()
	default:
		reply.Code = 404
		reply.Msg = "resource func err"
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
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
		go rpc.ServeConn(conn)
	}

}
