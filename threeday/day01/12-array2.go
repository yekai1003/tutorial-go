/*
   author:Yekai
   company:Pdj
   filename:12-array2.go
*/

package main

import (
	"fmt"
)

func main() {

	//Go语言当中的二维数组,可以理解为3行4列
	a2 := [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	//注意上述数组初始化的逗号
	fmt.Println(a2)
	//如何遍历该数组？可以写2层for循环搞定
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			fmt.Printf("i = %d, j = %d, val = %d\n", i, j, a2[i][j])
		}
	}
}
