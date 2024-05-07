package main

import "fmt"

func main() {
	fmt.Println("Iterating over array...")
	number := [5]int{1, 2, 3, 4, 5}
	for i := 0; i < len(number); i++ {
		fmt.Println(number[i])
	}

	fmt.Println("idiomatic iteration")
	numbers := [5]int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Println(index, value)
	}
}
