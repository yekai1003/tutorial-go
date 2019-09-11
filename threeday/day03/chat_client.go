/*
   author:Yekai
   company:Pdj
   filename:chat_client.go
*/
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

const (
	LOGIN = iota
	LOGOUT
	SETNAME
	BROADCAST
	PRIVATE
)

type ClientMsg struct {
	To      string  `json:"to"`
	MsgType int     `json:"msgtype"`
	Msg     string  `json:"msg"`
	DataLen uintptr `json:"datalen"`
}

func Help() {
	fmt.Println("1. setname:yourname")
	fmt.Println("2. all:yourmsg")
	fmt.Println("3. anyone:yourmsg")
}

func handle_conn(conn net.Conn) {
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Panic("Failed to Read ", err)
		}
		fmt.Println(string(buf[:n]))
		fmt.Printf("yekai's chat>")
	}
}

func main() {
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("Failed to Dial ", err)
		return
	}
	defer conn.Close()
	Help()
	go handle_conn(conn)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to yekai's chat")
	for {
		fmt.Printf("yekai's chat>")
		msg, err := reader.ReadString('\n')
		if err != nil {
			log.Panic("failed to ReadString ", err)
		}
		msg = strings.Trim(msg, "\r\n")
		if msg == "quit" {
			fmt.Println("byebye")
			break
		}
		msgs := strings.Split(msg, ":")
		msgtype := PRIVATE
		switch msgs[0] {
		case "set":
			msgtype = SETNAME
		case "all":
			msgtype = BROADCAST
		}
		climsg := ClientMsg{}
		climsg.DataLen = unsafe.Sizeof(climsg)
		climsg.MsgType = msgtype
		climsg.Msg = msgs[1]
		climsg.To = msgs[0]
		data, _ := json.Marshal(climsg)
		conn.Write(data)
	}
}
