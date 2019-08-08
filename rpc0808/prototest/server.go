package main

import (
	"errors"
	"log"
	"net"
	"net/rpc"
	"tutorial-go/rpc0808/goproto"
)

type MethodService struct{}

func (m *MethodService) Method(request *go0808.MethodParam, response *go0808.MethodReply) error {
	response.Code = 0
	response.Msg = "OK"
	switch request.GetMethod() {
	case "add":
		response.Reply = request.GetX() + request.GetY()
	case "sub":
		response.Reply = request.GetX() - request.GetY()
	case "mul":
		response.Reply = request.GetX() * request.GetY()
	case "div":
		if request.GetY() == 0 {
			response.Code = 100
			response.Msg = "Divisor is zero"
			return errors.New("params err")
		}
		response.Reply = request.GetX() / request.GetY()
	default:
		return errors.New("function 404")
	}
	return nil
}

func main() {
	rpc.RegisterName("MethodService", new(MethodService))

	listener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Panic("failed to Listen ", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panic("failed to Accept ", err)
		}
		go rpc.ServeConn(conn)
	}
}
