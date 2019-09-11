//编写http文件服务器，支持/Users/yekai/Downloads/webpath 目录
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

var index_html string = `<html><head><title>Index of ./</title></head>
<body bgcolor="#99cc99"><h4>Index of ./</h4>
<ul type=circle>
%s
</ul>
<address><a href="http://pdjedu.com/">xhttpd</a></address>
</body></html>`

func main() {
	//1. GET /hello.txt http/1.1 ==> hello.txt , /aa/bb/1.txt  ==> aa/bb/1.txt
	os.Chdir("/Users/yekai/Downloads/webpath") //切换工作目录
	//2. 侦听
	listener, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Panic("Failed to Listen", err)
	}
	defer listener.Close()
	for {
		//3. 循环等待新连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to Accept", err)
			continue //像小强一样活着
		}
		//4. 处理新连接的请求
		go handle_conn(conn)
	}

}

//处理连接
func handle_conn(conn net.Conn) {
	//1. 异常关闭连接
	defer conn.Close()
	reader := bufio.NewReader(conn)
	//定义正则表达式
	linereg := regexp.MustCompile("([a-zA-Z]+) (/.*) (.+)\r\n")
	//2. 循环等待消息
	for {
		//3. 获取请求行，解析请求行
		msg, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				fmt.Println("Failed to ReadString ", err)
				break
			}
		}
		//fmt.Println(msg)
		//GET /hello.txt HTTP/1.1
		msgs := linereg.FindAllStringSubmatch(msg, -1)
		//fmt.Println(msgs)
		if len(msgs) > 0 {
			// for k, v := range msgs[0] {
			// 	fmt.Println("k= ", k, ", v=", v)
			// }
			//处理请求资源
			go handle_method(conn, msgs[0][1], msgs[0][2])
		}
	}

}

func handle_method(conn net.Conn, method, url string) error {
	fmt.Println("method=", method, ", url=", url)
	if method == "GET" {
		//4. 如果资源不存在， 返回404错误
		// /hello.txt == > ./hello.txt
		path := "." + url
		info, err := os.Stat(path)
		if err != nil {
			fmt.Println("Failed to Stat", path, err)
			//有错误 404
			//处理404的逻辑
			errinfo, _ := os.Stat("error.html")
			//发送响应头（响应行，响应头，空行）
			sendContext(conn, getFileType("error.html"), "404", "NOT FOUND", errinfo.Size())
			//发送error.html
			sendFile(conn, "error.html")
			return err
		}
		//5. 如果资源存在：
		if info.IsDir() {
			//5.1 资源是一个目录，组织一个响应消息（html格式的文件）
			//形成超链接的列表项 <li><a href='xxx'>xxx.name</a></li>
			fd, err := os.Open(path)
			if err != nil {
				fmt.Println("Failed to Open Path", err)
				return err
			}
			infos, err := fd.Readdir(-1)
			if err != nil {
				fmt.Println("Failed to Readdir Path", err)
				return err
			}
			//拼接目录项 为列表项
			sendbuf := ""
			for _, v := range infos {
				//是否为目录
				if v.IsDir() {
					sendbuf += fmt.Sprintf("<li><a href='%s/'>%s</a></li>", v.Name(), v.Name())
				} else {
					sendbuf += fmt.Sprintf("<li><a href='%s'>%s</a></li>", v.Name(), v.Name())
				}
			}
			sendbuf = fmt.Sprintf(index_html, sendbuf)
			sendContext(conn, getFileType("error.html"), "200", "OK", 0)
			io.WriteString(conn, sendbuf)
		} else {
			//5.2 资源是一个文件, 组织发送文件的响应消息：响应头 + 正文
			//发送响应头（响应行，响应头，空行） + 正文
			sendContext(conn, getFileType(path), "200", "OK", info.Size())
			sendFile(conn, path)
		}

	}
	return nil
}

//组织http协议的
func sendContext(conn net.Conn, fileType, code, codeMsg string, length int64) {
	io.WriteString(conn, fmt.Sprintf("HTTP/1.1 %s %s\r\n", code, codeMsg)) //响应行
	if length > 0 {
		io.WriteString(conn, fmt.Sprintf("Content-Length: %d\r\n", length))
	}
	io.WriteString(conn, fmt.Sprintf("Content-Type: %s\r\n", fileType))
	io.WriteString(conn, "\r\n") //发送空行
}

//发送文件
func sendFile(conn net.Conn, filename string) error {
	// 读文件 - 写给网络
	fd, err := os.Open(filename)
	if err != nil {
		fmt.Println("Failed to Open ", filename, err)
		return err
	}
	defer fd.Close()
	n, err := io.Copy(conn, fd)
	fmt.Println("send file ok size = ", n)
	return err
}

//获取文件类型
func getFileType(filename string) string {
	fileType := "text/plain;charset=utf-8"
	if strings.HasSuffix(filename, ".html") {
		fileType = "text/html;charset=utf-8"
	} else if strings.HasSuffix(filename, ".jpg") {
		fileType = "image/jpeg"
	} else if strings.HasSuffix(filename, ".gif") {
		fileType = "image/jpeg"
	} else if strings.HasSuffix(filename, ".mp3") {
		fileType = "audio/mp3"
	}
	return fileType
}
