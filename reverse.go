package main

import "fmt"

type ListNode struct {
	Val  interface{}
	Next *ListNode
}

func Reverse(head *ListNode) *ListNode {

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

func Traverse(head *ListNode) {

	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}
func main() {
	one := ListNode{Val: 9}
	two := ListNode{Val: 4}
	three := ListNode{Val: 34}
	four := ListNode{Val: 98}

	one.Next = &two
	two.Next = &three
	three.Next = &four

	Traverse(&one)

	fmt.Println("reverse")

	rev := Reverse(&one)
	Traverse(rev)
}
