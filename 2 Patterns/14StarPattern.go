package main

import "fmt"

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}

//    *
//   ***
//  *****
// *******
func pattern(n int) {
	fmt.Println("Your patter is")
	for i := 0; i < n; i++ {
		for j := n - 1; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
}
