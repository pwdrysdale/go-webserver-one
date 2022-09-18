package main

import (
	"fmt"
	"math"
)

type square struct {
	length float64
}

type circle struct {
	radius float64
}

func (s square) area() float64 {
	return s.length * s.length
}

func (s square) circumf() float64 {
	return s.length * 4
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) circumf() float64 {
	return 2 * math.Pi * c.radius
}

// interfaces group types by methods they implement
type shape interface {
	area() float64
	circumf() float64
}

func printShapeInfo(s shape) {
	fmt.Printf("Area of %T is %0.2f \n", s, s.area())
	fmt.Printf("Circumference of %T is %0.2f \n", s, s.circumf())
}

func main() {
	s := []shape{
		square{length: 5},
		circle{radius: 5},
		square{length: 10},
	}
	for _, v := range s {
		printShapeInfo(v)
		fmt.Println("----------------")
	}
}
