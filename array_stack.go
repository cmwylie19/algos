package main

import "fmt"

type Stack []interface{}

func (s Stack) Peek() interface{} {
	if len(s) > 0 {
		return s[len(s)-1]
	} else {
		return "DONE"
	}

}
func (s *Stack) Push(v interface{}) {
	*s = append(*s, v)
}
func (s *Stack) Pop() interface{} {
	// last element
	last := (*s)[len(*s)-1]

	(*s) = (*s)[:len(*s)-1]

	return last
}
func main() {
	s := Stack{}
	s.Push("C")
	s.Push("a")
	s.Push("s")
	s.Push("E")
	s.Push("y")
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
}
