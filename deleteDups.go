package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Traverse(l *ListNode) {
	for l != nil {
		fmt.Println(l.Val)
		l = l.Next
	}
}
func swapPairs(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	prev := &ListNode{0, head}
	curr := prev
	for curr.Next != nil && curr.Next.Next != nil {
		a := curr.Next
		b := a.Next
		c := b.Next

		// b -> a
		b.Next = a

		// a -> c
		a.Next = c

		// b as first node
		curr.Next = b

		// iterate to the next two nodes
		curr = curr.Next.Next
	}

	// return everything after prev node
	return prev.Next
}

func reverseList(head *ListNode) *ListNode {

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
func deleteDuplicates(head *ListNode) *ListNode {

	// Represents the current node
	curr := head
	// |[a] -> a| -> a -> c -> d

	// Sliding window of two elements
	for curr != nil && curr.Next != nil {
		// if they are the same, skip b
		if curr.Next.Val == curr.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}

	return head

}

// a -> b -> c -> d -> e
func findMiddle(head *ListNode) *ListNode {
	curr, fast, slow := head, head, head
	for fast != nil && fast.Next != nil && fast.Next.Next != nil {
		fast = curr.Next.Next
		slow = curr.Next

	}

	return slow
}

func removeElements(head *ListNode, val int) *ListNode {
	prev := &ListNode{
		Val:  0,
		Next: nil,
	}
	prevPtr := prev

	// iterate through all nodes
	for head != nil {

		// remove all nodes with val == curr.Val
		if head.Val != val {
			prevPtr.Next = head
			prevPtr = prevPtr.Next
		}

		// iterate
		head = head.Next
	}
	return prev.Next
}

func main() {
	one := ListNode{
		Val: 6,
	}
	two := ListNode{Val: 23}
	three := ListNode{Val: 35}
	four := ListNode{Val: 6}
	five := ListNode{Val: 6}
	one.Next = &two
	two.Next = &three
	three.Next = &four
	four.Next = &five
	// Traverse(&one)
	// Traverse(&one)
	// Traverse(&one)

	testa := removeElements(&one, 6)
	Traverse(testa)

	// fmt.Println("Middle of LinkedList ", findMiddle(&one).Val)
	// Traverse(&one)
	// k := deleteDuplicates(&one)
	// Traverse(k)
	// reversed := reverseList(k)
	// Traverse(reversed)
}
