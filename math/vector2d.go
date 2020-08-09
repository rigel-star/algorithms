package main

import (
	"math"
)


type vector2d struct {
	x float64
	y float64
}


func (v *vector2d) add(v1 vector2d) {
	v.x = v.x + v1.x
	v.y = v.y + v1.y
}


func (v *vector2d) sub(v1 vector2d) {
	v.x = v.x - v1.x
	v.y = v.y - v1.y
}


func (v *vector2d) mul(v1 vector2d) {
	v.x = v.x * v1.x
	v.y = v.y * v1.y
}


func (v *vector2d) dist(v1 vector2d) float64 {
	dist := math.Sqrt(math.Pow(v1.x - v.x, 2) + math.Pow(v1.y - v.y, 2))
	return dist
}


func (v *vector2d) mag() float64 {
	return math.Sqrt(v.x * v.x + v.y * v.y)
}

//test
//func main() {
//	v := vector2d{4, 5}
//	v1 := vector2d{6, 3}
//
//	fmt.Print(v.dist(v1))
//}
