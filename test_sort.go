package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func createSortedLL(arr []int) *ListNode {

	head := &ListNode{}
	var prev *ListNode
	curr := head

	if len(arr) == 0 {
		return head
	}

	// create the first element
	if curr.Next == nil {
		nextNode := &ListNode{Val: arr[0]}
		head.Next = nextNode
		curr = curr.Next
	}

	//
	for i := 1; i < len(arr); i++ {
		if curr.Val < arr[i] {
			// must use prev
		}
	}

	return head.Next
}
func createLL(arr []int) *ListNode {
	head := &ListNode{Val: 0}
	curr := head

	for _, val := range arr {
		dummy := &ListNode{Val: val}
		curr.Next = dummy
		curr = curr.Next
	}

	return head.Next
}
func bubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < len(arr); i++ {
			if arr[i-1] > arr[i] {
				arr[i], arr[i-1] = arr[i-1], arr[i]
				swapped = true
			}
		}
	}

	return arr
}

func main() {
	arr := []int{9, 3, 2, 6, 89, 23, 97, 34, 1, 2}
	fmt.Println("Sorting")

	sorted := bubbleSort(arr)
	fmt.Printf("%+v", sorted)

	fmt.Println("\nLinked List")

	head := createLL(sorted)

	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}
