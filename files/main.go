package main

import (
	"io"
	"io/ioutil"
	"os"
)

func main() {
	println("Welcome to files in golang")

	content := "This needs to go in a file - google.com"
	file, err := os.Create("./myDocFile.txt")
	checknillErr(err)

	length, err := io.WriteString(file, content)

	checknillErr(err)

	println("Length is ", length)
	defer file.Close()

	readfile("./myDocFile.txt")

}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)

	checknillErr(err)

	println("Text data inside the file is  ", string(databyte))

}

func checknillErr(err error) {
	if err != nil {
		panic(err)
	}
}
