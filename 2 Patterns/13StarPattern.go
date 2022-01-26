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

//     1
//    12
//   123
//  1234
func pattern(n int) {
	for i := 1; i <= n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for p := 1; p <= i; p++ {
			fmt.Print(p)
		}
		fmt.Println("")
	}
}
