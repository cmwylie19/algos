package main

import "fmt"

func stringA(str string) string {
	if str == "AAAAAAAAAAA" {
		return ""
	}
	fmt.Println(str)
	stringA(str + "A")

	fmt.Printf("End of call where str = {%s}\n", str)
	return ""

}
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
	fmt.Println("")
	stringA("A")

}
