/*
   author:Yekai
   company:Pdj
   filename:19-polymorphic.go
*/
package main

import (
	"fmt"
)

type Animal interface {
	Sleeping()
	Eating()
}

type Cat struct {
	Color string
}

type Dog struct {
	Color string
}

//小猫结构体的方法
func (c Cat) Sleeping() {
	fmt.Println(c.Color, "Cat is sleeping")
}
func (c Cat) Eating() {
	fmt.Println(c.Color, "Cat is Eating")
}

//小狗结构体的方法
func (c Dog) Sleeping() {
	fmt.Println(c.Color, "Dog is sleeping")
}
func (c Dog) Eating() {
	fmt.Println(c.Color, "Dog is Eating")
}
func (c Dog) Print() {
	fmt.Println("Dog's color is", c.Color)
}

func Factory(color string, animal string) Animal {
	switch animal {
	case "dog":
		return &Dog{color}
	case "cat":
		return &Cat{color}
	default:
		return nil
	}
}

func main() {
	d1 := Factory("black", "dog")
	d1.Eating()
	c2 := Factory("white", "cat")
	c2.Sleeping()
}
