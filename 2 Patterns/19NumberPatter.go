package main

import "fmt"

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	patter(n)

}

// 1111
// 000
// 11
// 0

func patter(n int) {
	fmt.Println("Your patter is :")
	for i := n; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print((i + 1) % 2)
		}
		fmt.Println("")
	}
}
