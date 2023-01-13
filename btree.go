package main

import "fmt"

type Node struct {
	Val   interface{}
	Left  *Node
	Right *Node
}

func PreorderDFS(node *Node) {
	if node == nil {
		return
	}

	fmt.Println(node.Val)
	PreorderDFS(node.Left)
	PreorderDFS(node.Right)

}
func maxDepth(node *Node) int {
	if node == nil {
		return 0
	}

	return 1 + Max(maxDepth(node.Left), maxDepth(node.Right))
}
func Max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	root := &Node{
		Val: 0,
	}

	first := &Node{Val: "01"}
	second := &Node{Val: 2}
	three := &Node{Val: 3.0}
	f := &Node{Val: 99}
	ff := &Node{Val: 101}

	root.Left = first
	root.Right = second
	first.Left = three
	three.Left = f
	f.Left = ff

	PreorderDFS(root)
	fmt.Println("Max depth ", maxDepth(root))

}
