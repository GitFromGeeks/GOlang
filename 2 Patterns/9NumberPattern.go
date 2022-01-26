//n=4

// //n=4
// 1
// 1 1
// 1 2 1
// 1 2 2 1
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
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Print("1")
			} else {
				if j == 1 {
					fmt.Print("1")
				}
				fmt.Print("2")
			}

		}
		fmt.Println("")
	}
}
