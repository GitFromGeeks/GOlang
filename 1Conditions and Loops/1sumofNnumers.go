package main

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var sum = 0
	for i := 0; i < len(numbers); i++ {
		sum = sum + numbers[i]
	}
	println(sum)
}
