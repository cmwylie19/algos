package main

import (
	"fmt"
)

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func (head *ListNode) OnlyEvens() *ListNode {
	curr, evens := head, &ListNode{}
	curr_evens := evens
	for curr != nil {
		if curr.Val.(int)%2 == 0 {
			val := &ListNode{Val: curr.Val}
			curr_evens.Next = val
			curr_evens = curr_evens.Next
		}
		curr = curr.Next
	}

	return evens.Next
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

func Remove(head *ListNode, val int) *ListNode {
	var prev *ListNode

	curr := head

	for curr != nil {
		if curr.Val == val {
			if prev == nil {
				head = curr.Next
			} else {
				prev.Next = curr.Next
			}
		} else {
			prev = curr
		}
		curr = curr.Next

	}

	return head
}
func Traverse(head *ListNode) *ListNode {
	curr := head
	for curr != nil {
		fmt.Printf("%d ", curr.Val)
		curr = curr.Next
	}

	return head

}
func (head *ListNode) Reverse() *ListNode {
	var prev *ListNode
	curr := head
	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}

	return prev
}
func (head *ListNode) GetSum() int {
	sum, curr := 0, head

	for curr != nil {
		sum += curr.Val.(int)
		curr = curr.Next
	}

	return sum
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

	fmt.Println("\nOnly Evens")
	evens := Traverse(curr.OnlyEvens())
	fmt.Println("\nReverse")
	reversed_evens := Traverse(evens.Reverse())
	fmt.Println("\nGetSum: ")
	fmt.Printf("sum is %d\n", reversed_evens.GetSum())

	one := &ListNode{Val: 99}
	two := &ListNode{Val: 93}
	three := &ListNode{Val: 44}
	four := &ListNode{Val: 45}
	five := &ListNode{Val: 99}

	one.Next = two
	two.Next = three
	three.Next = four
	four.Next = five

	t := Remove(one, 99)
	Traverse(t)
	// fmt.Println("\nOnly Odds")
	// Traverse(curr.OnlyOdds())
}
