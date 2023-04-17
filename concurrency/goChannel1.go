package main

import (
	"fmt"
	"time"
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

func doSomeTask(val int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("out of the func with result %d", val)
}
