package main

import (
	"fmt"
	"time"
)

func main() {
	//result := doSomeTask(12)
	//fmt.Printf("%s\n", result)
	go doSomeTask(12)

	go func() {
		result := doSomeTask(12)
		fmt.Printf("%s\n", result) // will not print it

	}()
	fmt.Println("@ last line")
}

func doSomeTask(val int) string {
	time.Sleep(time.Second * 2)
	return fmt.Sprintf("out of the func with result %d", val)
}
