package main

import (
	"fmt"
)

func main() {
	numbers := []string{"a", "b", "c", "d"}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++")
	for index, number := range numbers {

		if index == 1 || index == 2 {
			fmt.Printf("%d -> %v \n", index, number)
		} else if index == 0 {
			continue
		} else {
			break
		}
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++")

	users := map[string]int{
		"daft": 1,
		"punk": 2,
	}
	for user, value := range users {
		fmt.Printf("%s --> %d\n", user, value)
	}

}
