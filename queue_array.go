package main

import "fmt"

var queue []interface{}

func Enqueue(q *[]interface{}, val interface{}) {
	(*q) = append(*q, val)
}
func Dequeue(q *[]interface{}) interface{} {
	if len(*q) == 0 {
		return "EMPTY"
	}
	top := (*q)[0]
	(*q) = (*q)[1:]
	return top
}

func Top(q *[]interface{}) interface{} {
	if len(*q) == 0 {
		return "EMPTY"
	}

	top := (*q)[0]
	return top
}
func main() {
	q := queue
	Enqueue(&q, "Up,")
	fmt.Println(Dequeue(&q))
	Enqueue(&q, "Hello ")
	Enqueue(&q, "World")
	Enqueue(&q, "Larry Hughes")

	for Top(&q) != "EMPTY" {
		fmt.Println(Dequeue(&q))
	}
	// q = append(q, "Hello ") // Enqueue
	// q = append(q, "world!")
	// for len(q) > 0 {
	// 	fmt.Print(q[0]) // First element
	// 	q = q[1:]       // Dequeue
	// }

}
