package main

import "fmt"

func Print(i int) {
	if i > 3 {
		return
	}

	fmt.Println(i)
	Print(i + 1)
	fmt.Printf("End of call where i = {%d}\n", i)
	return
}

func main() {
	Print(0)
}
