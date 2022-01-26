package main

import "fmt"

func main() {
	println("Check number is palindrome.")
	var num int
	fmt.Scanln(&num)
	println(checkPalindrome(num))

}

func checkPalindrome(num int) string {
	number := num
	isit := "NO"
	rev := 0
	for num > 0 {
		n := num % 10
		rev = rev*10 + n
		num = num / 10
	}
	if rev == number {
		isit = "Yes"
	}
	return isit
}
