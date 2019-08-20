/*
   author:Yekai
   company:Pdj
   filename:17-struct.go
*/
package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Sex   string
	Fight int
}

func main() {
	p1 := Person{"战五渣", 30, "man", 5}
	fmt.Printf("%+v\n", p1) // %+v 可以清晰的打印结构体数据
}
