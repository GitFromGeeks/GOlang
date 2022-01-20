package main

import "net/url"

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=gnind324sdf"

func main() {
	println("Welcome to URL handling in GoLang")
	println(myurl)
	//parsing

	result, _ := url.Parse(myurl)

	// println(result.Scheme)
	// println(result.Host)
	// println(result.Path)
	// println(result.Port())
	// println(result.RawQuery)
	qparams := result.Query()
	// println("The type of query params are : ", qparams)
	println(qparams["coursename"])

}
