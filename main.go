package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sidelength float64
}

//type area interface {
//get area() float64
//}
type shape interface {
	getArea() float64
}

func printArea(s shape) float64 {
	fmt.Println(s.getArea())
	return float64(s.getArea())
}

func main() {
	tr := triangle{}
	sq := square{}
	printArea(tr)
	printArea(sq)

}

func (t triangle) getArea() float64 {
	height := 5.5
	base := 6.2
	return 0.5 * height * base
}

func (s square) getArea() float64 {
	base := 6.1
	return base * base
}
