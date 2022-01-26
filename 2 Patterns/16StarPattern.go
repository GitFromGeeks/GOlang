package main

import "fmt"

func main() {
	fmt.Println("Enter any ODD number")
	var n int
	fmt.Scanln(&n)
	pattern(n)

}

func pattern(n int) {
	fmt.Println("Your pattern is :")
	patternHelper((n + 1) / 2)
}

func patternHelper(n int) {
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println("")
	}
	z := n
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(" ")
		}
		for k := 2*z + 1; k > 0; k-- {
			fmt.Print("*")
		}
		z--
		fmt.Println("")
	}
}

//     *
//    ***
//   *****
//  *******
// *********
//  *******
//   *****
//    ***
//     *

// Rohmbus shape
// for i := 0; i < n; i++ {
// 	for j := 0; j < i; j++ {
// 		fmt.Print(" ")
// 	}
// 	for k := 2*n + 1; k > 0; k-- {
// 		fmt.Print("*")
// 	}
// 	fmt.Println("")
// }
