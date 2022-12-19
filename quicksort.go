package main

import "fmt"

type Smallest struct {
	Index int
	Value int
}

func QuickSort(arr []int) []int {
	newArr := []int{}

	smallest := &Smallest{
		Index: 0,
		Value: arr[0],
	}

	for j := 0; j < len(arr); j++ {
		for i, v := range arr {
			if v < smallest.Value {
				smallest.Value = v
				smallest.Index = i
			}

		}

		newArr = append(newArr, smallest.Value)
		arr = append(arr[:smallest.Index-1], arr[smallest.Index:]...)

	}
	return newArr

}

var arr = []int{9, 5, 8, 4, 80, 1, 32, 43, 6, 8}

func main() {
	QuickSort(arr)
	fmt.Printf("%+v", arr)
}
