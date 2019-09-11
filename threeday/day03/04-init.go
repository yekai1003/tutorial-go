/*
   author:Yekai
   company:Pdj
   filename:04-init.go
*/
package main

import (
	"fmt"

	_ "github.com/yekai1003/tutorial/mathdemo"
	_ "github.com/yekai1003/tutorial/pkgdemo"
)

func init() {
	fmt.Println("main init is called")
}
func main() {

}
