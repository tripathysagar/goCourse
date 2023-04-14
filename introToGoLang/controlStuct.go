package main

import "fmt"

func main() {
	numbers := []string{"a", "b", "c", "d"}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(numbers[i])
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++")
	for index, number := range numbers {
		if index < 3 {
			fmt.Printf("%d -> %v \n", index, number)
		} else {
			break
		}
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++")

}
