package main

import "fmt"

func main() {
	resultch := make(chan string, 32)

	resultch <- "daft"
	resultch <- "punk"
	resultch <- "foo"
	resultch <- "bar"

	close(resultch) // we need to close the channel without it we will get deadlock error

	//for val := range resultch {
	//	fmt.Printf("result : %s\n", val)
	//
	//}

	select {

	default:
		val := <-resultch
		fmt.Printf("result : %s\n", val)

	}
}
