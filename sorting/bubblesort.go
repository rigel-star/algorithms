package main

const (
	arrSize = 6
)

//do test
var (
	arr = [arrSize]int{1, 5, 2, 9, 4, 3}
)

func sort(arr *[arrSize]int) {

	for i := 0; i < len(arr); i++ {
		for j := i; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

}

//create main func