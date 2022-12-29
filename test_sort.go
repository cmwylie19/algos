package main

import (
	"fmt"
)

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func (head *ListNode) OnlyOdds() *ListNode {
	curr, odds := head, &ListNode{}
	curr_odds := odds
	// Print message on odds
	for curr != nil {
		if curr.Val.(int)%2 != 0 {
			val := &ListNode{Val: curr.Val}
			curr_odds.Next = val
			curr_odds = curr_odds.Next
		}
		curr = curr.Next
	}
	return odds.Next
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

func Traverse(head *ListNode) {

	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
}

func main() {
	arr := []int{9, 3, 2, 6, 89, 23, 97, 34, 1, 2}
	fmt.Println("Sorting")

	sorted := bubbleSort(arr)
	fmt.Printf("%+v", sorted)

	fmt.Println("\nLinked List")

	head := createLL(sorted)
	curr := head
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}

	fmt.Println("\nOnly Odds")
	Traverse(curr.OnlyOdds())
}
