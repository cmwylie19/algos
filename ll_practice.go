package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Traverse(ll *ListNode) {
	for ll != nil {
		fmt.Println(ll.Val)
		ll = ll.Next
	}
}

func BubbleSort(arr []int) *ListNode {
	sigue := true

	for sigue {
		sigue = false

		for i := 0; i < len(arr); i++ {

		}
	}
}
func generateLL(arr []int) *ListNode {
	head := &ListNode{
		Val: 0,
	}
	curr := head
	for _, v := range arr {
		nextVal := &ListNode{
			Val: v,
		}
		curr.Next = nextVal
		curr = curr.Next
	}
	return head.Next
}

func main() {

	arr := []int{4, 6, 2, 1, 6, 8, 4, 2, 3, 5, 7, 9, 2, 12, 56, 9845, 234, 123}
	linkedList := generateLL(arr)
	Traverse(linkedList)

}
