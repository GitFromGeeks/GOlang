package main

import "fmt"

func main() {
	println("Sum of Even and Odd separetly")
	list := []int{}
	for i := 0; i < 4; i++ {
		v := 0
		fmt.Println("Enter value")
		fmt.Scanln(&v)
		list = append(list, v)
	}
	// for i := range list {
	// 	fmt.Println(list[i])
	// }
	fmt.Println(sumOfEvenOdd(list))
}

func sumOfEvenOdd(list []int) []int {
	even := 0
	odd := 0
	for i := range list {
		if list[i]%2 == 0 {
			even = even + list[i]
		} else {
			odd = odd + list[i]
		}
	}
	ls := []int{even, odd}
	return ls
}
