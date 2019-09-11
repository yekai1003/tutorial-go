package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"unsafe"
)

//服务器端与客户端传递消息格式：json
type ClientMsg struct {
	To      string  `json:"to"`      //接收者
	MsgType int     `json:"msgtype"` //消息类型
	Msg     string  `json:"msg"`     //实际消息
	Datalen uintptr `json:"datalen"` //消息长度 校验用
}

//消息类型
const (
	LOGIN = iota //iota 让枚举更优雅
	LOGOUT
	SETNAME
	BROADCAST
	PRIVATE
)

func Help() {
	fmt.Println("1. set:your name")
	fmt.Println("2. all:your msg")
	fmt.Println("3. anyone:your msg")
}

func handle_conn(conn net.Conn) {
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil || n <= 0 {
			log.Panic("Failed to Read", err)
		}
		fmt.Println(string(buf[:n]))
		fmt.Printf("pdj's chat>")
	}
}

func main() {
	Help()
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		log.Panic("Failed to Dial ", err)
	}
	defer conn.Close()
	//客户端的逻辑： 1. 读标准输入，发给服务器 2. 读服务器，响应在屏幕
	go handle_conn(conn)
	// 处理读标准输入-》发武器
	fmt.Println("Welcome to pdj's chat")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("pdj's chat>")
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Panic("Failed to ReadString ", err)
		}
		msg = strings.Trim(msg, "\r\n") //去回车换行
		if msg == "quit" {
			fmt.Println("Bye bye")
			break
		}
		msgs := strings.Split(msg, ":")
		if len((msgs)) == 2 {
			msgtype := PRIVATE
			switch msgs[0] {
			case "set":
				msgtype = SETNAME
			case "all":
				msgtype = BROADCAST
			}
			var climsg ClientMsg
			climsg.MsgType = msgtype
			climsg.To = msgs[0]
			climsg.Msg = msgs[1]
			climsg.Datalen = unsafe.Sizeof(climsg)
			data, _ := json.Marshal(climsg)
			conn.Write(data) //写给服务器
		}
	}
}
