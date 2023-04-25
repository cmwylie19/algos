package main

import (
	"fmt"
)

func BubbleSort(arr []int) []int {

	swapped := false
	for swapped {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				fmt.Println("%d is greater than %d", arr[i-1], 
arr[i])
				swapped = true
				arr[i-1], arr[i] = arr[i], arr[i-1]
			}
		}
	}

	return arr

}
func main() {
	nums := []int{3, 5, 1, 8, 34, 7, 9}
	num2s := BubbleSort(nums)
	for _, v := range num2s {
		fmt.Println(v)
	}
	//fmt.Printf("%+v", num2s)
}

