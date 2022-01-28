package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("Enter any number to check armstrong")
	var n string
	fmt.Scanln(&n)
	fmt.Println(armstrongChecker(n))
}

func armstrongChecker(n string) bool {
	number, _ := strconv.ParseInt(n, 10, 64)
	nm := number
	var sum int64
	for i := 0; i < len(n); i++ {
		num := number % 10
		sum = sum + num*num*num
		number = number / 10
	}
	if sum == nm {
		return true
	}
	return false
}
