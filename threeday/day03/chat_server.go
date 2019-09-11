/*
   author:Yekai
   company:Pdj
   filename:chat_server.go
*/
package main

import (
	"encoding/json"
	"fmt"
	"net"
	"unsafe"
)

/*
	需求：
		1. 支持多个客户端
		2. 客户端有登陆和退出工作
		3. 客户端可以发送广播消息或私信消息
	设计：
		1. 广播消息

*/

type ChatMsg struct {
	From string
	To   string
	Msg  string
}

type ClientMsg struct {
	To      string  `json:"to"`
	MsgType int     `json:"msgtype"`
	Msg     string  `json:"msg"`
	DataLen uintptr `json:"datalen"`
}

var (
	chan_login     chan ChatMsg
	chan_logout    chan ChatMsg
	chan_msgcenter chan ChatMsg
)

const (
	LOGIN = iota
	LOGOUT
	SETNAME
	BROADCAST
	PRIVATE
)

type Client struct {
	conn net.Conn
	name string
}

var mapClients map[string]Client
var clientKeys map[string]string

func main() {

	//通道初始化
	chan_login = make(chan ChatMsg)
	chan_logout = make(chan ChatMsg)
	chan_msgcenter = make(chan ChatMsg)
	//map 初始化
	mapClients = make(map[string]Client)
	clientKeys = make(map[string]string)

	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		fmt.Println("Failed to Listen", err)
		return
	}

	go msg_center()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept", err)
			continue
		}
		go handle_conn(conn)
	}
}

func logout(conn net.Conn, from string) {
	conn.Close()
	msg := ChatMsg{from, "all", from + ":" + "logout"}
	chan_logout <- msg
	delete(mapClients, from)
}

func handle_conn(conn net.Conn) {
	from := conn.RemoteAddr().String()
	fmt.Println(from, "login")
	defer logout(conn, from)
	//连接就算登陆
	msg := ChatMsg{from, "all", from + ":" + "login"}
	chan_login <- msg
	cli := Client{conn, ""}
	clientKeys[from] = from
	mapClients[from] = cli

	//提供服务
	buf := make([]byte, 256)
	for {

		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to Read data")
			return
		}
		msg.Msg = string(buf[:n])
		var climsg ClientMsg
		err = json.Unmarshal(buf[:n], &climsg)
		if err != nil {
			fmt.Println("failed to Unmarshal ", err)
			continue
		}
		if climsg.DataLen != unsafe.Sizeof(climsg) {
			fmt.Println("msg format err", err)
			continue
		}
		//chan_broadcast <- msg
		sendmsg := ChatMsg{conn.RemoteAddr().String(), "all", climsg.Msg}
		switch climsg.MsgType {
		case SETNAME:
			//直接修改名字
			from := conn.RemoteAddr().String()
			mapClients[from].name = climsg.Msg
			//mapClients[mapClients[from].name] = mapClients[from]
			clientKeys[climsg.Msg] = from
			sendmsg.Msg = from + " set name=" + climsg.Msg + " sucess"
			sendmsg.From = "server"
		case BROADCAST:
		case PRIVATE:
			sendmsg.To = climsg.To
		}

		chan_msgcenter <- sendmsg
	}
}

//消息中心
func msg_center() {
	for {
		select {
		case msg := <-chan_login:
			msgsend(msg)
		case msg := <-chan_logout:
			msgsend(msg)
		case msg := <-chan_msgcenter:
			fmt.Println(msg)
			msgsend(msg)
		}
	}
}

func msgsend(msg ChatMsg) {
	data, _ := json.Marshal(msg)
	if msg.To == "all" {
		//需要广播
		for _, v := range mapClients {
			if v.conn.RemoteAddr().String() != msg.From {
				v.conn.Write(data)
			}

		}
	} else {
		fmt.Println(msg)
		key, ok := clientKeys[msg.To]
		if !ok {
			fmt.Println("User not exists")
			return
		}
		cli, ok := mapClients[key]
		if !ok {
			fmt.Println("User not exists mapClients")
			return
		}
		cli.conn.Write(data)
	}
}
