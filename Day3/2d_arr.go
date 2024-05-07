package main

import "fmt"

func main() {
	fmt.Println("Iterating over 2d Array")
	var matrix [3][3]int

	count := 1
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			matrix[i][j] = count
			count++
		}
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
		fmt.Println()
	}
}
