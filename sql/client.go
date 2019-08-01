/*

	1. 连接到数据库
	2. 循环等待sql输入
	3. 判断一下sql 是查询类，还是非查询类
	3.1 如果没有结果集，执行Exec
	3.2 如果有结果集，执行Query，遍历结果集，打印结果集

*/
package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Client struct {
	connstr string
	driver  string
	dbconn  *sql.DB
}

// 入口函数-构造Client结构体
func NewClient(user, pass, dbname, protocol string) *Client {
	connstr := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", user, pass, protocol, dbname)
	return &Client{connstr, "mysql", nil}
}

//连接数据库
func (cli *Client) Conn() error {
	db, err := sql.Open(cli.driver, cli.connstr)
	if err != nil {
		fmt.Println("faild to open mysql ", err)
		return err
	}
	cli.dbconn = db
	return cli.dbconn.Ping()
}

//客户端运行
func (cli *Client) Run() {
	//循环读取sql语句- 标准输入
	rd := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("yekai-mysql>")
		sqlstr, err := rd.ReadString('\n')
		if err != nil {
			log.Panic("failed to ReadString ", err)
		}
		//处理一下接收的字符串
		sqlstr = strings.Trim(sqlstr, "\r\n")
		//fmt.Printf("sql==[%s]\n", sqlstr)
		sqls := []byte(sqlstr)

		if len(sqls) > 6 {
			//判断是有结果集还是没有结果集
			//1. select show desc
			if string(sqls[:6]) == "select" || string(sqls[:4]) == "show" || string(sqls[:4]) == "desc" {
				//有结果集
				cli.sql_query(sqlstr)
			} else {
				//无结果集
				cli.sql_exec(sqlstr)
			}
		}
		if sqlstr == "quit" {
			fmt.Println("bye bye")
			break
		}
	}
}

func (cli *Client) sql_exec(sqlstr string) {
	//执行 insert_sql
	result, err := dbconn.Exec(sqlstr)
	if err != nil {
		fmt.Println("failed to sql_exec ", err)
	}
	rowsaffed, _ := result.RowsAffected()
	fmt.Println(rowsaffed)
}

func (cli *Client) sql_query(sqlstr string) {
	//执行 query
	rows, err := dbconn.Query(sqlstr)
	if err != nil {
		fmt.Println("failed to Query ", err)
		return
	}

	//打印表头
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("failed to Columns ", err)
		return
	}
	for _, v := range cols {
		fmt.Printf("%s\t", v)
	}
	fmt.Println("\n--------------------------------") //华丽的分割线
	//打印数据
	//列的个数？
	colCount := len(cols)

	//values := make([]string, colCount)
	values := make([]sql.NullString, colCount)
	oneRows := make([]interface{}, colCount) //Scan 使用接收
	for k, _ := range values {
		oneRows[k] = &values[k] //结果集内存地址绑定
	}
	for rows.Next() {
		err = rows.Scan(oneRows...)
		if err != nil {
			fmt.Println("failed to Scan ", err)
			break
		}
		for _, v := range values {
			//如果 valid=false代表该字段是null
			//fmt.Printf("%v\t", v)
			if v.Valid {
				fmt.Printf("%s\t", v.String)
			} else {
				fmt.Printf("NULL\t")
			}
		}
		fmt.Println()
	}

}
