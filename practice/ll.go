package main

import "fmt"

type Node struct {
	Val  interface{}
	Next *Node
}

func (head *Node) Populate(arr []int) {
	curr := head

	// get to the end of the ll
	for curr.Next != nil {
		curr = curr.Next
	}
	for _, val := range arr {
		temp := &Node{Val: val}
		curr.Next = temp
		curr = curr.Next
	}

}
func (head *Node) Traverse() {
	curr := head

	for curr != nil {
		fmt.Println(curr.Val)
		curr = curr.Next
	}
}
func RecursiveTraverse(head *Node) {
	if head == nil {
		fmt.Println("DONE")
	} else {
		fmt.Print(head.Val, " ")
		RecursiveTraverse(head.Next)
	}
}
func Count(start *int, head *Node) int {
	fmt.Println("Called ", start)
	if head.Next != nil {
		*start += 1
		Count(start, head.Next)
	}
	return *start + 1

}
func (head *Node) Add(node *Node) {
	curr := head

	// add to the end

	for curr.Next != nil {
		curr = curr.Next
	}
	curr.Next = node
}
func (head *Node) Remove() {
	curr := head

	// find the second to last
	for curr.Next.Next != nil {
		curr = curr.Next
	}
	curr.Next = nil
}
func main() {
	node := &Node{Val: 1}

	//arr := []int{34, 5, 7, 3, 6, 8, 4, 2, 678, 9}
	// node.Populate(arr)
	// node.Populate(arr)
	node.Add(&Node{Val: 2})
	node.Add(&Node{Val: 3})
	node.Add(&Node{Val: 4})
	node.Traverse()

	fmt.Println("Recursive Traverse")
	t := &node
	RecursiveTraverse(*t)
	fmt.Println("Count")
	va := 0
	fmt.Println(Count(&va, *t))
	// node.Remove()
	// node.Remove()

	// fmt.Println("LinkedList")
	// node.Traverse()
}
