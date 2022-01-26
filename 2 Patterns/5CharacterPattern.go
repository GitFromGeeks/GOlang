// A
// B C
// C D E
// E F G H
package main

import "fmt"

//TODO: pattern is not done yet!

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}
func pattern(n int) {
	ch := 65
	fmt.Println("Your pattern is")
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Printf("%c", rune(ch+j))
		}
		fmt.Println("")
	}
}
