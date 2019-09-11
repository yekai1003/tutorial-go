/*
   author:Yekai
   company:Pdj
   filename:01-tcp_server.go
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	//1. 指定为ipv4协议，绑定IP和端口，启动侦听
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Panic("Failed to Listen", err) //输出错误信息，并且执行panic错误
	}
	defer listener.Close() //收尾工作
	for {
		//2. 循环等待新连接到来，Accept阻塞等待状态
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept ", err)
			continue
		}
		fmt.Println("New conn->", conn.RemoteAddr().String())
		//3. 启动匿名函数来处理
		go func(conn net.Conn) {
			defer conn.Close() //收尾工作
			buf := make([]byte, 256)
			for {
				//从客户端读数据
				n, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF { //这种错误代表客户端先关闭了，属于正常范围内的错误
						fmt.Println("Client ", conn.RemoteAddr().String(), " is Closed")
						break
					}
					fmt.Println("Failed to Read data ", err)
					break
				}
				//回射服务器，收到什么，写回什么
				n, err = conn.Write(buf[:n])
				if err != nil {
					fmt.Println("Failed to Write to client ", conn.RemoteAddr().String(), err)
					break
				}
			}

		}(conn)
	}
}
