package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	patter(n)
}
func patter(n int) {
	fmt.Println("Your patter is :")
	z := n
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 1; k <= z; k++ {
			fmt.Print(i + k)
		}
		fmt.Println("")
		z--
	}
	p := n - 2
	for i := 1; i <= n-1; i++ {
		for j := p; j > 0; j-- {
			fmt.Print(" ")
		}
		for k := i; k >= 0; k-- {
			fmt.Print(i + p - k + 1)
		}
		fmt.Println("")
		p--
	}
}

// 123456
//  23456
//   3456
//    456
//     56
//      6
//     56
//    456
//   3456
//  23456
// 123456
