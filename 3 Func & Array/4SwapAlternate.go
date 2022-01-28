package main

import "fmt"

func main() {
	fmt.Println("Enter an array of numbers")
	list := []int{}
	for i := 0; i < 4; i++ {
		n := 0
		fmt.Scanln(&n)
		list = append(list, n)
	}
	newList := []int{}
	for i := 0; i < len(list); i += 2 {
		newList = append(newList, list[i+1])
		newList = append(newList, list[i])
	}
	fmt.Println(newList)

}
