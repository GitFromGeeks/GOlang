//n=4
// 1
// 1 1
// 2 0 2
// 3 0 0 3
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
			fmt.Print(" ", i)
		}
		fmt.Println("")
	}
}
