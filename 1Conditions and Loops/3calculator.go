package main

import "fmt"

func main() {
	var number int
	println("1 Addition")
	println("2 Subtration")
	println("3 Multiplication")
	println("4 Division")
	fmt.Scanln(&number)
	switch number {
	case 1:
		var num1 int8
		var num2 int8
		println("Addition Selected")
		println("First Number")
		fmt.Scanln(&num1)
		println("Second Number")
		fmt.Scanln(&num2)
		println("Addition", num1+num2)
	case 2:
		println("Subtraction Selected")
		var num1 int8
		var num2 int8
		println("First Number")
		fmt.Scanln(&num1)
		println("Second Number")
		fmt.Scanln(&num2)
		println("Subtraction", num1-num2)
	case 3:
		println("Multiplication Selected")
		var num1 int8
		var num2 int8
		println("First Number")
		fmt.Scanln(&num1)
		println("Second Number")
		fmt.Scanln(&num2)
		println("Multiplication", num1*num2)
	case 4:
		println("Division Selected")
		var num1 int8
		var num2 int8
		println("First Number")
		fmt.Scanln(&num1)
		println("Second Number")
		fmt.Scanln(&num2)
		println("Division", num1/num2)
	default:
		println("Invalid Operation")

	}
}
