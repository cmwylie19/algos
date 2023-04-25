package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func preorderdfs(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	preorderdfs(node.Left)
	preorderdfs(node.Right)
}
func minDepth(root *Node) int {

	if root == nil {
		return 0
	}

	left := minDepth(root.Left)

	right := minDepth(root.Right)

	fmt.Println("Left is ", left, " Right is ", right, " root.Val ", root.Val)

	// if left or right node is empty
	if left == 0 || right == 0 {
		fmt.Println("Hits this case")
		return left + right + 1
	}

	return min(left, right) + 1

}

func main() {
	fmt.Println("Hi")

	zero := &Node{
		Val: 0,
	}
	one := &Node{
		Val: 1,
	}
	two := &Node{
		Val: 2,
	}
	three := &Node{
		Val: 3,
	}
	four := &Node{
		Val: 4,
	}
	five := &Node{
		Val: 5,
	}
	six := &Node{
		Val: 6,
	}

	zero.Left = one
	zero.Right = two
	one.Left = three
	one.Right = four
	four.Right = six

	two.Right = five

	preorderdfs(zero)
}
