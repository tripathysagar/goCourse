package main

import (
	"fmt"
)

func main() {
	resultch := make(chan string) // unbuffered channel
	fmt.Println(resultch)
	/*
		resultch <- "foo" // is full now  -> it will BLOCK -> BLOCK here
		result := <-resultch

		fmt.Println(result)
	*/

	go func() {
		result := <-resultch
		fmt.Println(result)
	}()
	resultch <- "foo"
}
