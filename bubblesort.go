package main

import "fmt"

func BubbleSort(arr []int) []int {

	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i] < arr[i-1] {
				swapped = true
				arr[i], arr[i-1] = arr[i-1], arr[i]

			}
		}
	}
	return arr
}

func main() {
	arr := []int{33, 1, 34, 89, 34, 1, 9, 2}
	arr = BubbleSort(arr)
	for _, v := range arr {
		fmt.Println(v)
	}
}
