package main

import (
	"fmt"
)

func main() {
	println("Fahreinheit to Celcious Converter")
	var Fahreinheit int
	print("In Fahreinheit :")
	fmt.Scanln(&Fahreinheit)
	println("In Celcious :- ", (Fahreinheit-32)*5/9)
}
