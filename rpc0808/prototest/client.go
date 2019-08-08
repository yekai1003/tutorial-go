package main

import (
	"fmt"
	"log"
	"net/rpc"
	"tutorial-go/rpc0808/goproto"
)

func main() {
	cli, err := rpc.Dial("tcp", "localhost:9999")
	if err != nil {
		log.Panic("failed to Dial ", err)
	}
	defer cli.Close()

	var param = &go0808.MethodParam{
		Method: "add",
		X:      10,
		Y:      20,
	}

	var reply = &go0808.MethodReply{}

	if cli.Call("MethodService.Method", param, reply) != nil {
		log.Panic("failed to Call Method ", err)
	}

	fmt.Println(reply)

	param.Method = "mul"

	if cli.Call("MethodService.Method", param, reply) != nil {
		log.Panic("failed to Call Method ", err)
	}

	fmt.Println(reply)

	param.Y = 0
	param.Method = "div"
	if err = cli.Call("MethodService.Method", param, reply); err != nil {
		log.Panic("failed to Call Method ", err)
	}

	fmt.Println(reply)
}
