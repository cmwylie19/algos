package main

import "fmt"

func max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// s = "eceba", 2
func longestUniqueSubString(s string, k int) int {

	kv := make(map[string]int)
	left, m := 0, 0

	if len(s) == 0 {
		return m
	}

	for right := 0; right < len(s); right++ {
		kv[string(s[right])]++

		for len(kv) > k {
			fmt.Println("Constraint violated: ", len(kv))
			kv[string(s[left])]--

			if kv[string(s[left])] == 0 {
				delete(kv, string(s[left]))
			}
			left++

		}

		// check max
		m = max(m, right-left+1)

	}
	return m

}

// var kv = map[string]int{
// 	"hi":   1,
// 	"u":    2,
// 	"have": 3,
// 	"a":    4,
// 	"day":  5,
// }

func main() {
	fmt.Println(longestUniqueSubString("eceba", 2))
	// delete(kv, "hi")
	// fmt.Println("length", len(kv))
	// fmt.Printf("%+v\n", kv)
}
