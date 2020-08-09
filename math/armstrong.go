package main

/*
Armstrong number is a number that is equal to the sum of the cubes of its digits.
e.g. 153 = 1 ^ 3 + 5 ^ 3 + 3 ^ 3
*/

func isArmstrong(n int) bool {
	temp := n
	remainder := 0
	sum := 0

	for n > 0 {
		remainder = n % 10
		sum += remainder * remainder * remainder
		n /= 10
	}

	return temp == sum
}


//test
//func main() {
//	fmt.Println(isArmstrong(417))
//}