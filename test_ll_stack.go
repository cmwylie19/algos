package main

import "fmt"

type Stack []interface{}

func (stack *Stack) Push(v interface{}) {
	*stack = append(*stack, v)
}
func (stack *Stack) Pop() interface{} {
	if len(*stack) == 0 {
		return "EMPTY"
	}
	latest := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return latest
}
func (stack *Stack) Peek() interface{} {
	if len(*stack) == 0 {
		return "EMPTY"
	}

	return (*stack)[len(*stack)-1]
}
func main() {
	one := &Stack{}
	one.Push("shit")
	one.Push("ass")
	one.Push("Robot")
	for one.Peek() != "EMPTY" {
		fmt.Printf("\n %s", one.Pop())
	}
}
