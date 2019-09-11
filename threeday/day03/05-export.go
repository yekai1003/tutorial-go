/*
   author:Yekai
   company:Pdj
   filename:05-export.go
*/
package main

import (
	"fmt"

	"github.com/yekai1003/tutorial/pkgdemo"
)

func init() {
	fmt.Println("main init is called")
}
func main() {
	//这样会报错，sex字段非导出，不能填写
	p1 := pkgdemo.ExternalPerson{"yekai", 40, "man"}
	fmt.Println(p1)
	//这样没有问题，NewPerson是可以导出的
	p2 := pkgdemo.NewPerson("fuhongxue", 37, "man")
	fmt.Println(p2)

}
