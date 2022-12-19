package main

import (
	"fmt"
)

var romanNumeralDict = map[int]string{
	1000: "M",
	500:  "D",
	100:  "C",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func main() {
	fmt.Println("j")
	s := New()
	s.Push("2")
	s.Push(2)
	s.Push("Casey")
	s.Push(999)
	fmt.Println(s.Len())
	for s.Len() > 0 {
		fmt.Println(s.Pop())
	}
}
