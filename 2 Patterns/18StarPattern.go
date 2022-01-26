package main

import "fmt"

func main() {
	fmt.Println("Enter any number")
	var n int
	fmt.Scanln(&n)
	pattern(n)
}

// *
//  * *
//    * * *
//      * * * *
//        * * * * *
//          * * * * * *
//        * * * * *
//      * * * *
//    * * *
//  * *
// *
func pattern(n int) {
	fmt.Println("Your patter is :")
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print("  ")
		}
		for k := 0; k <= i; k++ {
			fmt.Print(" *")
		}
		fmt.Println("")
	}
	for i := n - 1; i >= 0; i-- {
		for j := i; j > 1; j-- {
			fmt.Print("  ")
		}
		for k := i; k > 0; k-- {
			fmt.Print(" *")
		}
		fmt.Println("")
	}
}
