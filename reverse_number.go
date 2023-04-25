package main

import "fmt"

// Remove the last digit from our number and append it to the end of our reversed number var until our original number is gone
func ReverseNumber(num int) int {
	reversed := 0

	for num > 0 {
		// isolate last digit in the number
		lastDigit := num % 10

		// multiply by 10 so that ones becomes tens, tens become hundreds
		// append lastDigit into end of reverse
		reversed = reversed*10 + lastDigit

		// remove the last digit
		num /= 10

	}
	return reversed
}
func main() {
	fmt.Println("Reverse 1234", ReverseNumber(1234))
}
