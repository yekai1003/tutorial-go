/*
   author:Yekai
   company:Pdj
   filename:19-inherit.go
*/
package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Sex   string
	Fight int
}

func (p *Person) setAge(age int) {
	p.Age = age
}

type SuperMan struct {
	Strength int
	Speed    int
	Person
}

func (s SuperMan) Print() {
	fmt.Printf("Name = %s, Age = %d, Sex = %s, Fight = %d\n", s.Name, s.Age, s.Sex, s.Fight)
	fmt.Printf("strength = %d, fight = %d\n", s.Strength, s.Speed)
}

func main() {
	s1 := SuperMan{
		Strength: 100,
		Speed:    99,
		Person: Person{
			Name:  "Kelak",
			Age:   40,
			Sex:   "man",
			Fight: 5000,
		},
	}
	fmt.Println(s1)
	s1.setAge(41)
	s1.Print()
}
