package main

import (
	"fmt"
)

//定义业务类型
const (
	login = iota // iota = 0
	logout
	user    = iota + 2 //iota = 2,user = 2+2 = 4
	account = iota * 3 //iota = 3, account = 3*3 = 9
)

const (
	apple, banana = iota + 1, iota + 2 // iota = 0
	peach, pear                        //iota = 1
	orange, mango                      //iota = 2
)

func main() {
	fmt.Println(login, logout, user, account)
	fmt.Println(apple, banana, peach, pear, orange, mango)
}
