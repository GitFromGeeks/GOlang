package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("On Topic of Channel in GOlang")

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	//recieve only
	go func(ch <-chan int, wg *sync.WaitGroup) {
		fmt.Println(<-mych)
		wg.Done()
	}(mych, wg)
	//send only
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		close(mych)
		wg.Done()
	}(mych, wg)

	wg.Wait()

}
