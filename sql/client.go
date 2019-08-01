/*
   file    : client.go
   author  : yekai
   company : pdj(pdjedu.com)
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

func NewClient(user, pass, dbname, protocol string) *Client {
	connstr := fmt.Sprintf("%s:%s@%s/%s?charset=utf8", user, pass, protocol, dbname)
	return &Client{connstr, "mysql", nil}
}

func (cli *Client) Conn() error {
	db, err := sql.Open(cli.driver, cli.connstr)
	if err != nil {
		fmt.Println("failed to open database ", err)
		return err
	}
	cli.dbconn = db
	return cli.dbconn.Ping()
}

func (cli *Client) Run() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to mysql client ")
	for {
		fmt.Printf("yekai-mysql>")
		sqlstr, err := reader.ReadString('\n')
		if err != nil {
			log.Panic("failed to ReadString ", err)
		}
		sqlstr = strings.Trim(sqlstr, "\r\n")
		sqls := []byte(sqlstr)
		if len(sqls) > 6 {
			if string(sqls[:6]) == "select" || string(sqls[:4]) == "show" || string(sqls[:4]) == "desc" {
				//result set sql
				cli.query_sql(sqlstr)
			} else {
				//no result set sql
				cli.exec_sql(sqlstr)
			}
		}
		if sqlstr == "quit" {
			fmt.Println("bye bye ")
			break
		}
	}
}

func (cli *Client) exec_sql(xsql string) {
	result, err := cli.dbconn.Exec(xsql)
	if err != nil {
		fmt.Println("failed to exec sql:", xsql, err)
		return
	}
	rowsaff, _ := result.RowsAffected()
	fmt.Println("RowsAffected:", rowsaff)
}

func (cli *Client) query_sql(xsql string) {
	rows, err := cli.dbconn.Query(xsql)
	if err != nil {
		fmt.Println("failed for query sql:", err)
		return
	}
	cols, err := rows.Columns()
	if err != nil {
		fmt.Println("failed for get columns:", err)
		return
	}
	colCount := len(cols)

	//fmt.Println(colCount, "\n", cols)
	for _, v := range cols {
		fmt.Printf("%s\t", v)
	}
	fmt.Println("\n----------------------------------------")

	//使用NullString可以很好的支持空值
	values := make([]sql.NullString, colCount)
	oneRows := make([]interface{}, colCount)
	for k, _ := range values {
		oneRows[k] = &values[k] //将查询结果的返回地址绑定，这样才能变参获取数据
	}

	for rows.Next() {

		//扫描结果集，一定要在Next调用后，方可使用
		err = rows.Scan(oneRows...)
		if err != nil {
			fmt.Println("failed to Scan result set", err)
			break
		}
		//fmt.Println(values)
		for _, v := range values {
			if v.Valid {
				fmt.Printf("%s\t", v.String)
			} else {
				fmt.Printf("%s\t", "NULL")
			}
		}
		fmt.Println()
	}
}
