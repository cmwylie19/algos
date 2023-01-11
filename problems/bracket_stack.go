package main

import (
	"fmt"
)

type Stack []interface{}

func (s *Stack) Push(val interface{}) {
	(*s) = append(*s, val)

}
func (s Stack) Peek() interface{} {
	if len(s) > 0 {
		return s[len(s)-1]
	} else {
		return "DONE"
	}

}
func (s *Stack) Pop() interface{} {
	// last element in the array
	latest := (*s)[len(*s)-1]

	// remove last element from array
	(*s) = (*s)[:len(*s)-1]

	return latest

}

func matchingBrackets(str string) bool {

	var s = &Stack{}
	// iterate over the string
	for _, val := range str {
		if string(val) == "(" {
			s.Push(string(val))
		} else if string(val) == "{" {
			s.Push(string(val))
		} else if string(val) == "[" {
			s.Push(string(val))
		} else if string(val) == ")" {
			if s.Peek() == "DONE" || fmt.Sprintf("%s", s.Pop()) != "(" {
				return false
			}
		} else if string(val) == "}" {
			if s.Peek() == "DONE" || fmt.Sprintf("%s", s.Pop()) != "{" {
				return false
			}
		} else if string(val) == "]" {
			if s.Peek() == "DONE" || fmt.Sprintf("%s", s.Pop()) != "[" {
				return false
			}
		}
	}

	if s.Peek() == "DONE" {
		return true
	} else {
		return false
	}

}

func main() {

	// name := "CaseyWylie"
	// age := 34

	s1 := "({})"
	s2 := "(){}[]"
	s3 := "(]"
	s4 := "({)}"
	s5 := "(("
	s6 := "{[]}"
	fmt.Printf("%s  %t\n", s1, matchingBrackets(s1))
	fmt.Printf("%s  %t\n", s2, matchingBrackets(s2))
	fmt.Printf("%s  %t\n", s3, matchingBrackets(s3))
	fmt.Printf("%s  %t\n", s4, matchingBrackets(s4))
	fmt.Printf("%s  %t\n", s5, matchingBrackets(s5))
	fmt.Printf("%s  %t\n", s6, matchingBrackets(s6))

}
