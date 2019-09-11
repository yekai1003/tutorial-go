package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	var v1 int
	var x, y = 123, "hello"
	a, str := 456, "world"
	fmt.Println(v1, x, y, a, str)
	//定义复数，注意i一定要与数字紧紧相连，否则会被当成字符i处理
	var c1 complex64 = 4 + 3i
	fmt.Println(c1)

	var b *int = &a
	fmt.Println(*b)
	*b = 100
	fmt.Println(a, *b)

	swap(x, a)
	fmt.Println(x, a)
	swap2(&x, &a)
	fmt.Println(x, a)
}

func swap(a, b int) {
	temp := a
	a = b
	b = temp
}

func swap2(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}
