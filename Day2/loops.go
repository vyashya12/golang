package main

import "fmt"

func main() {
	fmt.Println("Basic for loop")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("For loop as a while loop...")
	i := 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("for loop with range (for slices, arrays, maps, strings)")
	nums := []int{1, 2, 3, 4, 5}
	for index, value := range nums {
		fmt.Printf("Index: %d,Value: %d\n", index, value)
	}

	fmt.Println("for loop with a range (for strings)")
	for index, runeValue := range "hello world" {
		fmt.Printf("Index: %d,Value: %c\n", index, runeValue)
	}
}
