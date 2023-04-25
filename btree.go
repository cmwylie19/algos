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
func maxDepthIterate(node *Node) int {
	copyNode := node
	left, right := 1, 1
	for node.Left != nil {
		node = node.Left
		left += 1
	}

	for copyNode.Right != nil {
		copyNode = copyNode.Right
		right += 1
	}

	if right > left {
		return right
	}

	return left
}
func maxDepth(node *Node) int {
	if node == nil {
		return 0
	}
	left := maxDepth(node.Left)
	right := maxDepth(node.Right)
	fmt.Println("Val is ", node.Val)

	return 1 + Max(left, right)
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
	four := &Node{Val: 90210}
	f1 := &Node{Val: 0}
	f2 := &Node{Val: 0}
	f3 := &Node{Val: 0}
	f4 := &Node{Val: 0}
	f5 := &Node{Val: 0}
	f6 := &Node{Val: 0}
	f7 := &Node{Val: 99}
	f := &Node{Val: 99}
	ff := &Node{Val: 101}

	root.Left = first
	root.Right = second
	first.Left = three
	three.Left = f
	three.Right = four
	four.Left = f1
	f1.Left = f2
	f2.Left = f3
	f3.Left = f4
	f4.Left = f5
	f5.Left = f6
	f6.Left = f7
	f.Left = ff

	PreorderDFS(root)
	fmt.Println("Max depth ", maxDepth(root))

}
