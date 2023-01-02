package main

import "fmt"

func Traverse(head *node) {
	for head != nil {
		fmt.Println(head.value)
		head = head.next
	}
}

type (
	LinkedList struct {
		head *node
		tail *node
		size int
	}
	node struct {
		value interface{}
		next  *node
	}
)

func Middle(head *node) interface{} {
	slow := head
	fast := head
	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}
	return slow.value
}

func main() {

	one := &node{value: 1}
	two := &node{value: 2}
	three := &node{value: 3}
	one.next = two
	two.next = three

	Traverse(one)
	// print(Middle(one))

	// ll := &LinkedList{head: one, tail: three, size: 3}

	// for ll.head != nil {
	// 	fmt.Println(ll.head.value)
	// 	ll.head = ll.head.next
	// }

	// fmt.Printf("%+v", ll)

}

func print(s interface{}) {
	fmt.Println(s)
}
