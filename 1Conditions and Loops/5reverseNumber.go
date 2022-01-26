package main

import "fmt"

func main() {
	println("Enter any number to Reverse")
	var number int
	fmt.Scanln(&number)
	rev := reverseNumber(number)
	println("reverse is :", rev)
}

func reverseNumber(num int) int {
	rev := 0
	for num > 0 {
		n := num % 10
		rev = rev*10 + n
		num = num / 10
	}
	return rev
}
