// *
// * *
// * * *
// * * * *
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
	fmt.Println("Your Patter is")
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}

}
