package main

import (
	"fmt"
	"log"
	"net/rpc"
	"tutorial-go/rpc/proto"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}

	var reply = &goprotoc.MsgValue{}
	var param = &goprotoc.MsgValue{
		Value: "yekai",
	}

	err = client.Call("HelloService.Hello", &param, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply)

	var reply1 = &goprotoc.MethodVal{}
	var param1 = &goprotoc.MethodVal{
		Method: "mul",
		X:      10,
		Y:      23,
	}

	err = client.Call("HelloService.Method", &param1, &reply1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply1)

}
