// 4444
// 333
// 22
// 1

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}
func pattern(n int) {
	for i := n; i > 0; i-- {
		for j := i; j > 0; j-- {
			fmt.Print(i)
		}
		fmt.Println("")
	}
}
