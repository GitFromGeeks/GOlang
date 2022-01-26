// //n=3

// A
// BB
// CCC

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
	ch := 64
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf(" %c", rune(ch+i))
		}
		fmt.Println("")
	}
}
