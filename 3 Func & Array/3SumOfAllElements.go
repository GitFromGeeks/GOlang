package main

import "fmt"

func main() {
	fmt.Println("Enter some numbers")
	list := []int{}
	for i := 0; i < 4; i++ {
		n := 0
		fmt.Scanln(&n)
		list = append(list, n)
	}
	sum := 0
	for i := range list {
		sum = sum + list[i]
	}
	fmt.Println("Sum of all elements is :", sum)
}
