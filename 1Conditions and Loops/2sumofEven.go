package main

func main() {
	println("Sum of even numbers")
	numbers := []int{2, 3, 4, 5, 6, 7}
	var sum = 0
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			sum = sum + numbers[i]
		}
	}
	println(sum)
}
