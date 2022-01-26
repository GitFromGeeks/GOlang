package main

import "fmt"

//TODO: Pattern not yet done
func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}

// 1
// 232
// 34543
// 4567654
func pattern(n int) {
	fmt.Println("Your pattern is :")
	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			if i == j {
				fmt.Print(j)
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
