package main

import "fmt"

func main() {
	fmt.Println("Enter an array of numbers")
	list := []int{}
	for i := 0; i < 5; i++ {
		n := 0
		fmt.Scanln(&n)
		list = append(list, n)
	}
	newList := []int{}
	fmt.Println(newList)
}

// 1 2 2 3 4 4
// out 1 3
