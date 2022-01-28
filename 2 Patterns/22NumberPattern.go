package main

import "fmt"

func main() {
	fmt.Println("Enter any EVEN number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}

// 1  2  3  4
// 9 10 11 12
// 13 14 15 16
// 5  6  7  8

func pattern(n int) {
	fmt.Println("Your patter is :")
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 {
				fmt.Print(" ", j)
			} else if i == 4 {
				fmt.Print(" ", i+j)
			} else {
				fmt.Print(" ", j+4*i)

			}
		}
		fmt.Println("")
	}
}
