// 1
// 2 1
// 3 2 1
// 4 3 2 1
package main

import "fmt"

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}

func pattern(n int) {
	fmt.Println("your pattern is")
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(i - j)
		}
		fmt.Println("")
	}
}
