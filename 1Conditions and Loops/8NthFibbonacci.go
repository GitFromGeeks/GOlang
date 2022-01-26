package main

import "fmt"

func main() {
	fmt.Println("Nth Fibbonacci")
	var n int
	fmt.Println("Enter any number")
	fmt.Scanln(&n)
	fmt.Println("Your fabbonacci is : ", NthFibbo(n))

}

// 0 1 1 2 3 5 8

func NthFibbo(n int) int {
	fibs := []int{0, 1}
	for i := 1; i <= n; i++ {
		a := fibs[i] + fibs[i-1]
		fibs = append(fibs, a)
	}
	return fibs[n-1]
}
