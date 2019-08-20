/*
   author:Yekai
   company:Pdj
   filename:18-enclosure.go
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

//第一种方法，传入2个点
func distance(p1, p2 Point) float64 {
	return math.Sqrt((p2.x-p1.x)*(p2.x-p1.x) + (p2.y-p1.y)*(p2.y-p1.y))
}

//第二种方法，以一个点作为参照点，再传入一个点
func (this Point) distance2(p Point) float64 {
	return math.Sqrt((this.x-p.x)*(this.x-p.x) + (this.y-p.y)*(this.y-p.y))
}

func main() {
	p1 := Point{0.0, 0.0}
	p2 := Point{3.0, 4.0}
	fmt.Println(p2, "between", p1, "distance is ", distance(p1, p2))
	fmt.Println(p2.distance2(p1))
}
