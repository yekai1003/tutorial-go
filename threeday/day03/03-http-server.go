/*
   author:Yekai
   company:Pdj
   filename:03-http-server.go
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

//router hello
func HelloUserServer(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	fmt.Println(path)
	users := strings.Split(path, "/")
	fmt.Println(len(users), users)
	if len(users) == 3 {
		io.WriteString(w, users[1]+","+users[2])
	}

}

//router byebye
func ByeUserServer(w http.ResponseWriter, req *http.Request) {

	path := req.URL.Path
	users := strings.Split(path, "/")
	if len(users) == 3 {
		io.WriteString(w, users[1]+","+users[2])
	}

}

func main() {
	//设置hello的路由
	http.HandleFunc("/hello/", HelloUserServer)
	//设置byebye的路由
	http.HandleFunc("/bye/", ByeUserServer)
	//侦听并提供web服务，所有的事情都在前面设置的路由函数实现了
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
