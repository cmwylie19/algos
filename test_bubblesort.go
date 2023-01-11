package main

import "fmt"

func asc(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		// 6 2 5 8 3 4 67 25 4 1
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				swapped = true
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
	}
	return arr
}

func main() {
	arr := []int{6, 3, 23, 75, 1, 23, 8}
	ascending := asc(arr)
	// descending := desc(ascending)
	fmt.Printf("Ascending:\n%+v\n", ascending)

	// fmt.Printf("Descending:\n%+v\n", descending)
}
