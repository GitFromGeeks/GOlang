package main

func main() {
	defer println("At very last of the program")
	defer println("One")
	defer println("Two")
	println("Hello world")
	defer myDefer()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer println(i)
	}
}
