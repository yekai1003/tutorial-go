/*
   author:Yekai
   company:Pdj
   filename:01-json.go
*/
package main

import (
	"fmt"
)

/*
{
	"Name":"yekai",
	"Age":35,
	"Sex":"man",
	"Favirate":{
		"wuqi":"feidao",
		"food":"huasheng"
	}
}
*/

func main() {
	a := 10

	if a > 10 { //左括号一定要写在表达式同行，与函数要求相同
		fmt.Println("My God ,a is ", a)
	} else if a < 10 {
		fmt.Println("a is too small")
	} else {
		fmt.Println("a == ", a)
	}

}
