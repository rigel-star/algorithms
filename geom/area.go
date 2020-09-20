package main

import (
	"fmt"
	"math"
)

const (
	PI = math.Pi
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rect struct {
	width, height float64
}

func ( c *circle ) circumference() float64 {
	return 2 * PI * c.radius
}

func ( c circle ) area() float64 {
	return PI * ( c.radius * c.radius )
}

func ( r rect ) area() float64 {
	return r.width * r.height
}

func measure( s shape ) {
	fmt.Println( "Area: ", s.area() )
}

func main() {
	a := 3e+8
	c := circle{5}
	measure( c )

	fmt.Println( a )
}