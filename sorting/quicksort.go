package main

import (
	"fmt"
)

/*Source: https://en.wikipedia.org/wiki/Quicksort#Lomuto_partition_scheme*/

const (
	arrSize = 6
)

var (
	arr = [arrSize]int{2, 5, 3, 8, 1, 9}
	itr = 0
)

func sort(arr *[arrSize]int, low int, high int) {

	if low < high {
		p := partition(arr, low, high)

		sort(arr, low, p - 1)
		sort(arr, p + 1, high)
	}
}

/*partioning array into sub arrays*/
func partition(arr *[arrSize]int, low int, high int) int {

	pivot := arr[high]
	i := low

	for j := low; j <= high; j++ {
		if arr[j] < pivot {
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp

			i += 1
		}

		itr += 1
	}

	tmp2 := arr[i]
	arr[i] = arr[high]
	arr[high] = tmp2

	return i
}


/*test*/
func main() {
	fmt.Println(arr)

	sort(&arr, 0, len(arr) - 1)

	fmt.Println(arr)
	fmt.Println("Total iterations: ", itr)

}