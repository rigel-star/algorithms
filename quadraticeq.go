package main

import (
	"fmt"
	"math"
)

type equation struct {
	a float64
	b float64
	c float64
}

func (eq *equation) solve() (float64, float64) {
	formPos := -eq.b + math.Sqrt(eq.b - (4 * eq.a * eq.c))
	formNeg := -eq.b - math.Sqrt(eq.b - (4 * eq.a * eq.c))

	return formPos / 2 * eq.a, formNeg / 2 * eq.a
}

//test
func main() {

	eq := equation{3, 2, 1}
	res1, res2 := eq.solve()

	fmt.Println(res1, res2)
}
