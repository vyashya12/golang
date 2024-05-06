package main

import "fmt"

func main() {
	x := 10
	fmt.Println("Basic IF")
	if x > 5 {
		fmt.Println("x is bigger than 5")
	}

	fmt.Println("If else")

	x = 4
	if x > 5 {
		fmt.Println("x is bigger than 5")
	} else {
		fmt.Println("x is not bigger than 5")
	}

	x = 10
	fmt.Println("if else if else")
	if x > 10 {
		fmt.Println("x is bigger than 10")
	} else if x == 10 {
		fmt.Println("x is exadtly 10")
	} else {
		fmt.Println("x is less than 10")
	}

	fmt.Println("With Initialization Statement")
	if y := 20; y > 10 {
		fmt.Println("y is gretaer than 10")
	}
	// y is not accessible outside of the block
}
