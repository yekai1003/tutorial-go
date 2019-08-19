package main

import "fmt"

//import "unsafe"

func main1() {
	const LENGTH int = 10
	const WIDTH = 5
	const a, b, c = 1, false, "str" //多重赋值

	//注意用 := ,代表定义变量area
	area := LENGTH * WIDTH
	//Printf与Printfln的区别是格式化以及自带换行
	fmt.Printf("area is %d\n", area)
	println(a, b, c)
}

func test() int {
	return 10
}

const (
	a = "abc"
	b = test()
	c = len(a)
)

func main() {
	println(a, b, c)
}
