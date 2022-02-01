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
	fmt.Println("Your list of numbers", list)
	fmt.Println(findDuplicate(list))
}

func findDuplicate(list []int) []int {
	newList := []int{}
	for i := range list {
		x := list[i] % len(list)
		list[x] = list[x] + len(list)
	}

	return newList
}

// 1 1 2 4 5 6 6
// out 1 6
