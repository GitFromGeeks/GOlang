// n=4

// 1
// 1 1
// 1 1 1
// 1 1 1 1
package main

import "fmt"

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}
func pattern(n int) {
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ", 1)
		}
		fmt.Println("")
	}
}
