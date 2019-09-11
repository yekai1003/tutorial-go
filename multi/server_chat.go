/*
	并发聊天室需求：
	 	1. 支持多个客户端
		2. 广播或私信
		3. 上线通知
		4. 下线通知
		5. 设置昵称
	分析：
		1. 网络编程 net + goroutine
		2. 广播： 客户端自己发给其他客户端吗？流程：客户端A发给服务器端，服务器端广播给其他客户端 ，服务器端必须记录 连接列表
		3. 同步机制，多个goroutine需要传递消息 channel 设置
		4. 消息的类型： 广播或私信 消息结构体体现 设计一个channel
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"unsafe"
)

//消息类型
const (
	LOGIN = iota //iota 让枚举更优雅
	LOGOUT
	SETNAME
	BROADCAST
	PRIVATE
)

type ChatMsg struct {
	From, To, Msg string
}

//服务器端与客户端传递消息格式：json
type ClientMsg struct {
	To      string  `json:"to"`      //接收者
	MsgType int     `json:"msgtype"` //消息类型
	Msg     string  `json:"msg"`     //实际消息
	Datalen uintptr `json:"datalen"` //消息长度 校验用
}

//消息中心的通道
var chan_msgcenter chan ChatMsg

//定义连接列表
var mapClients map[string]net.Conn //key->conn
var keyClients map[string]string   // name ->key ; key->conn

func main() {

	//初始化
	chan_msgcenter = make(chan ChatMsg)
	mapClients = make(map[string]net.Conn)
	keyClients = make(map[string]string)

	// 设置侦听
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		log.Panic("Failed to Listen ", err)
	}
	//延迟关闭
	defer listener.Close()
	//启动消息中心处理
	go msg_center()
	//循环接收各个客户端连接
	for {
		//接收客户端连接请求
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept ", err)
			continue
		}
		//处理各个客户端的请求 goroutine
		go handle_conn(conn)
	}
}

//断开连接处理
func logout(conn net.Conn, from string) {
	//关闭连接
	defer conn.Close()

	//组织一个消息发给消息中心，广播的形式
	msg := ChatMsg{from, "all", from + "->logout"}
	chan_msgcenter <- msg
	delete(mapClients, from)
}

// 处理请求的函数
func handle_conn(conn net.Conn) {
	//4. 善后工作
	//1. 拿到客户端信息
	from := conn.RemoteAddr().String() //ip:port
	defer logout(conn, from)

	//2. 放到map中
	mapClients[from] = conn
	//2.1 上线通知
	msg := ChatMsg{from, "all", from + "->login"}
	chan_msgcenter <- msg
	//3. 处理读写行为
	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to Read data ", err)
			break
		}
		//处理客户端的消息请求
		var climsg ClientMsg
		err = json.Unmarshal(buf[:n], &climsg)
		if err != nil {
			fmt.Println("Failed to Unmarshal ", err)
			continue //死缓
		}
		if climsg.Datalen != unsafe.Sizeof(climsg) {
			fmt.Println("msg format err ", climsg)
			continue //死缓
		}
		fmt.Println(climsg)
		//组织一个消息
		from := conn.RemoteAddr().String()
		chatmsg := ChatMsg{from, "all", climsg.Msg}
		switch climsg.MsgType {
		case SETNAME:
			//1.名字存在哪 2. 名字怎么用(聊天的时候用) - > 通过名字可以找到连接信息
			//fmt.Println("setname------", msg)
			keyClients[climsg.Msg] = from
			chatmsg.Msg = from + " set name=" + climsg.Msg + " success"
			chatmsg.From = "server" //发送者改为服务器
		case BROADCAST:
		case PRIVATE:
			chatmsg.To = climsg.To
		}
		//发送到消息中心
		chan_msgcenter <- chatmsg

	}

}

//消息处理中心
func msg_center() {
	for {
		msg := <-chan_msgcenter //读消息中心的channel
		//发送：广播或私信
		fmt.Println(msg)
		go send_msg(msg)
	}
}

func send_msg(msg ChatMsg) {
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Failed to Marshal ", err)
		return
	}
	if msg.To == "all" {
		//广播
		//遍历map 发送
		for _, v := range mapClients {
			//不给自己发
			if msg.From != v.RemoteAddr().String() {
				v.Write(data)
			}
		}
	} else {
		//私信
		from, ok := keyClients[msg.To]
		if !ok {
			fmt.Println("User not exists ", msg.To)
			return
		}
		conn, ok := mapClients[from]
		if !ok {
			fmt.Println("client not exists ", from)
			return
		}
		conn.Write(data)
	}
}
