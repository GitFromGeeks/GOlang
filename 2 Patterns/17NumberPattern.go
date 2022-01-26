package main

import "fmt"

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}

// 1        1
// 12      21
// 123    321
// 1234  4321
// 1234554321

func pattern(n int) {
	fmt.Println("Your pattern is :")
	z := n + 1
	for i := 0; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for k := 0; k < 2*z-2; k++ {
			fmt.Print(" ")
		}
		for l := i; l >= 1; l-- {
			fmt.Print(l)
		}
		fmt.Println("")
		z--
	}
}
