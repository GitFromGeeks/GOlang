package main

import "fmt"

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}

// 33333
// 32223
// 32123
// 32223
// 33333

func pattern(n int) {
	fmt.Println("your patter is :")
	for i := 0; i < 2*n-1; i++ {
		for j := 0; j < 2*n-1; j++ {
			fmt.Print(n)
		}
		fmt.Println("")
	}
}
