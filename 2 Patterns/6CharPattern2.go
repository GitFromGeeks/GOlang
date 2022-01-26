// n=5
// E
// D E
// C D E
// B C D E
// A B C D E

//TODO: pattern is not done yet
package main

import "fmt"

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
		for j := 0; j <= i; j++ {
			fmt.Printf(" %c", rune(ch-j+n))
		}
		fmt.Println(" ")
	}
}
