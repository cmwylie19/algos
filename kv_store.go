package main

import "fmt"

var one = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
}

var two = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
}

const FOUR = 4

var four = map[string]interface{}{
	"one":   1,
	"two":   "two",
	"three": 3.34,
	"four":  FOUR,
}

var three = map[string][]interface{}{
	"one":   {1, 2, 3},
	"two":   {4, 5, 6},
	"three": {7, 8, 9},
}

func KeyLength3(m map[string]interface{}) int {
	return len(m)
}
func KeyLength2(m map[string][]interface{}) int {
	return len(m)
}
func KeyLength(m map[string]int) int {
	return len(m)
}
func main() {
	fmt.Println(KeyLength(one))
	fmt.Println(KeyLength(two))
	fmt.Println(KeyLength2(three))
	fmt.Println(KeyLength3(four))

	fmt.Println("j")
}
