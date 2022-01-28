package main

import "fmt"

func main() {
	fmt.Println("Enter a number to chech Fibonacci")
	var n int
	fmt.Scanln(&n)
	fmt.Println(FibChecker(n))

}

func FibChecker(n int) bool {
	first := 0
	second := 1
	num := 0
	if n == 0 || n == 1 {
		return true
	} else {
		for i := 1; i < n; i++ {
			num = first + second
			first = second
			second = num
			if num == n {
				return true
			}
		}

	}
	return false
}
