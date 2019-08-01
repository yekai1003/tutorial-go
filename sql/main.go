package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

var dbconn *sql.DB

func init() {
	db, err := sql.Open("mysql", "root:abc123@tcp(10.211.55.3:3306)/yekai?charset=utf8")
	if err != nil {
		log.Panic("failed to open mysql ", err)
	}
	if err = db.Ping(); err != nil {
		log.Panic("failed to ping mysql ", err)
	}
	dbconn = db
}

var (
	drop_table   string = "drop table person"
	create_table string = "create table person(name varchar(30), age int)"
	insert_table string = "insert into person(name, age) values(?,?)"
)

func sql_test() {
	result, err := dbconn.Exec(drop_table)
	if err != nil {
		log.Panic("failed to drop table ", err)
	}
	fmt.Println(result.LastInsertId())

	result, err = dbconn.Exec(create_table)
	if err != nil {
		log.Panic("failed to create table ", err)
	}
	fmt.Println(result.LastInsertId())

	result, err = dbconn.Exec(insert_table, "yekai", "40")
	if err != nil {
		log.Panic("failed to insert table ", err)
	}
	fmt.Println(result.LastInsertId())
}

func main() {
	imort_file("./hzw.txt")
	// cli := NewClient("root", "abc123", "yekai", "tcp(10.211.55.3:3306)")
	// if cli.Conn() != nil {
	// 	log.Panic("failed to conn to mysql ")
	// }
	// //cli.query_sql("show tables")
	// cli.Run()
}

func insert_line(line string) {
	s := strings.Fields(line)
	_, err := dbconn.Exec(insert_table, s[0], s[1])
	if err != nil {
		log.Panic("failed to insert into table ", err)
	}
}

func imort_file(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		log.Panic("failed to open file ", err)
	}
	rd := bufio.NewReader(f)
	// line, prefix, err := rd.ReadLine()
	// defer f.Close()
	// fmt.Println(string(line), prefix)
	// array := strings.Fields(string(line))
	// fmt.Println(array, array[0], "---", array[1])
	for {
		line, _, err := rd.ReadLine()
		if err != nil && err != io.EOF {
			log.Panic("failed to readline ", err)
		}
		if err == io.EOF {
			fmt.Println("read over")
			break
		}
		insert_line(string(line))
	}
}
