package main

import "fmt"

type Queue []interface{}

func (queue *Queue) Enqueue(i interface{}) {
	*queue = append(*queue, i)
}
func (queue *Queue) Dequeue() interface{} {
	if len(*queue) == 0 {
		return "EMPTY"
	}
	front := (*queue)[0]
	*queue = (*queue)[1:]
	return front
}

func main() {
	q := &Queue{}
	q.Enqueue(1)
	q.Enqueue("hi")
	q.Enqueue("Sup")
	q.Enqueue("Enyza")

	for len(*q) != 0 {
		fmt.Println(q.Dequeue())
	}

}
