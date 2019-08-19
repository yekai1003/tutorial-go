/*
   author:Yekai
   company:Pdj
   filename:06-switch.go
*/
package main

import (
	"fmt"
)

func main() {
	var fruit string
	fmt.Println("Please input a fruit's name:")
	fmt.Scanf("%s", &fruit)
	switch fruit {
	case "banana":
		fmt.Println("I want 2 banana!")
	case "orange":
		fmt.Println("I want 3 orange!")
	case "apple":
		fmt.Println("I want an apple!")
	case "pear":
		fmt.Println("I do not like pear!")
	default:
		fmt.Println("Are you kidding me?")
	}
}
