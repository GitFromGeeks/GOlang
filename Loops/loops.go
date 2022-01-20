package main

func main() {
	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	for index, day := range days {
		println("at index %v, the day is %v", index, day)
	}

}
