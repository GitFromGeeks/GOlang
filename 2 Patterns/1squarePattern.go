//4 4 4 4
//4 4 4 4
//4 4 4 4
//4 4 4 4

package main

import "fmt"

func main() {
	fmt.Println("Pattern")
	var n int
	fmt.Println("Enter any number")
	fmt.Scanln(&n)
	squarePattern(n)

}

func squarePattern(n int) {
	fmt.Println("Your Patter is :")
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(n)
		}
		fmt.Println("")
	}
}
