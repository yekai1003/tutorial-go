package animal

import (
	"fmt"
)

type Cat struct {
	Name  string
	Color string
	Age   uint
}

func NewCat(name, color string, age uint) *Cat {
	return &Cat{name, color, age}
}

func (c *Cat) Sleeping() {
	fmt.Println(c.Color, "Cat", c.Name, "is sleeping")
}

func (c *Cat) Eating() {
	fmt.Println(c.Color, "Cat", c.Name, "is Eating")
}

func (c *Cat) Print() {
	fmt.Printf("+%v\n", c)
}
