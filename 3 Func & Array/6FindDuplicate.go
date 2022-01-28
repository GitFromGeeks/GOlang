package main

import "fmt"

func main() {
	fmt.Println("enter an array of numbers")
	list := []int{}
	for i := 0; i < 5; i++ {
		n := 0
		fmt.Scanln(&n)
		list = append(list, n)
	}
	newList := []int{}
	fmt.Println(newList)
}

// 1 1 2 4 5 6 6
// out 1 6
