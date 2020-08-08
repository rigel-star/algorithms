package main

import (
	"fmt"
	"sort"
)

var (
	arr = []int{2, 5, 4, 10, 6, 8, 9, 3, 1}
)

func search(arr []int, val int, low int, high int) int  {
	mid := low + (high - low) / 2

	if val >= low {
		if val == arr[mid] {
			return mid
		}

		if arr[mid] > val {
			return search(arr, val, low, mid - 1)
		}
		return search(arr, val , mid + 1, high)
	}
	return -1
}

func main() {
	sort.Ints(arr)

	index := search(arr, 10, 0, len(arr) - 1)

	fmt.Println(arr)

	if index != -1 {
		fmt.Println("Found at index: ", index)
	}
}
