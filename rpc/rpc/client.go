package main

import (
	"fmt"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8888")
	defer client.Close()

	var reply string
	if client.Call("HelloService.Hello", "yekai", &reply) != nil {
		log.Panic("failed to rpc-call ", err)
	}

	fmt.Println("rpccall retutn:", reply)
}
