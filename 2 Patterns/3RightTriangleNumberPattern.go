// 1
// 2 2
// 3 3 3
package main

import "fmt"

func main() {
	fmt.Println("Enter any random number")
	var n int
	fmt.Scanln(&n)
	pattern(n)

}

func pattern(n int) {
	fmt.Println("Your patter is")
	for i := 0; i <= n; i++ {
		for j := 0; j < i; j++ {
			fmt.Print(i)
		}
		fmt.Println("")
	}
}
