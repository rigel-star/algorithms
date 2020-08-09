package main

import (
	"math"
)

func isPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n < 2 || n % 2 == 0 {
		return false
	}

	lim := math.Sqrt(float64(n))

	for i := 3; i <= int(lim); i += 2 {
		if n % i == 0 {
			return false
		}
	}

	return true
}

//test
//func main() {
//	fmt.Println(isPrime(11)) //will print true
//	fmt.Println(isPrime(25)) //will print false
//	fmt.Println(isPrime(53)) //will print true
//}
