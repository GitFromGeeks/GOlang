//n=4
// 1
// 1 1
// 2 0 2
// 3 0 0 3
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
				fmt.Print(i)
			} else {
				if j == 1 {
					fmt.Print(i)
				}
				fmt.Print("0")
			}

		}
		fmt.Println("")
	}
}
