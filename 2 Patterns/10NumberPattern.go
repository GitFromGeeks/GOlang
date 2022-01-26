// n=4

//  1234
//  123
//  12
//  1

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
		for j := 1; j <= i; j++ {
			print(j)
		}
		fmt.Println("")
	}
}
