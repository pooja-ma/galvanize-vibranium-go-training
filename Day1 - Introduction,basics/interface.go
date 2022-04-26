package main

import (
	"fmt"
	"math"
)

// An interface type is defined as a set of method signatures.
// A value of interface type can hold any value that implements those methods.

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type square struct {
	width float64
}

func (r rect) area() float64 {
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim(

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	//s := square{width: 6}
	measure(r)
	measure(c)
	//measure(s)

}
