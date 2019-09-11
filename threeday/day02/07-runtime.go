/*
   author:Yekai
   company:Pdj
   filename:07-runtime.go
*/
package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	fmt.Println("Go Max proc =", runtime.GOMAXPROCS(0))
	runtime.GOMAXPROCS(2)
	fmt.Println("Go Max proc =", runtime.GOMAXPROCS(0))

}
