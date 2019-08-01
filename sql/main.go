package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var (
	drop_sql   string = "drop table person"
	create_sql string = "create table person(name varchar(30), age int)"
	insert_sql string = "insert into person(name, age) values(?,?)"
)

var dbconn *sql.DB

//init函数会自动加载，而且运行一次
func init() {
	db, err := sql.Open("mysql", "root:abc123@tcp(10.211.55.3)/yekai?charset=utf8")
	if err != nil {
		log.Panic("failed to Open mysql ", err)
	}

	//测试网络是否通畅,ping可以定位连接错误
	if db.Ping() != nil {
		log.Panic("failed to ping mysql ", err)
	}
	fmt.Println("connect to mysql ok")
	dbconn = db
}

func main2() {
	fmt.Println("hello world")
	//端口默认是 3306
	//打开mysql

	// 关闭mysql连接
	defer dbconn.Close()

	query_exec("select * from person")
	return

	//执行 drop_sql
	result, err := dbconn.Exec(drop_sql)
	if err != nil {
		log.Panic("failed to drop table ", err)
	}
	fmt.Println(result.RowsAffected())

	//执行 create_sql
	result, err = dbconn.Exec(create_sql)
	if err != nil {
		log.Panic("failed to create_sql ", err)
	}
	fmt.Println(result.RowsAffected())

	//执行 insert_sql
	result, err = dbconn.Exec(insert_sql, "yekai", 35)
	if err != nil {
		log.Panic("failed to insert_sql ", err)
	}
	fmt.Println(result.RowsAffected())

	//导入文件
	import_file("./hzw.txt")
}

func insert_exec(name, age string) {
	//执行 insert_sql
	_, err := dbconn.Exec(insert_sql, name, age)
	if err != nil {
		log.Panic("failed to insert_sql ", err)
	}
}

//需求：将hzw.txt文件内容读取，插入到数据库中 shell

func import_file(filename string) {
	// 读取文件
	f, err := os.Open(filename)
	if err != nil {
		log.Panic("failed to Open file ", filename, err)
	}
	defer f.Close()
	// 循环读取每一条
	rd := bufio.NewReader(f)
	// line, prefix, err := rd.ReadLine()
	// fmt.Println(string(line), prefix, err)
	for {
		line, _, err := rd.ReadLine()
		//有错误，并且不是文件末尾
		if err != nil && err != io.EOF {
			log.Panic("failed to readline ", err)
		}
		//如果读到文件末尾结束
		if err == io.EOF {
			break
		}

		//分割字符串
		fields := strings.Fields(string(line))
		fmt.Println(fields)
		// 调用exec
		insert_exec(fields[0], fields[1])
	}

}

func query_exec(rsql string) {
	//查询第一步：Query 得到rows
	rows, err := dbconn.Query(rsql)
	if err != nil {
		log.Panic("failed to Query ", err)
	}
	//rows是结果集
	fmt.Println(rows.Columns())

	//第二步：遍历结果集
	for rows.Next() {
		var Name string
		var Age int
		//Scan函数必须在Next执行返回true之后才可以使用
		err = rows.Scan(&Name, &Age)
		if err != nil {
			log.Panic("failed to Scan ", err)
		}
		fmt.Println("name=", Name, ",age=", Age)
	}
}

func main() {
	cli := NewClient("root", "abc123", "yekai", "tcp(10.211.55.3)")
	if cli.Conn() != nil {
		log.Panic("failed to connect to mysql ")
	}

	cli.Run()
}
