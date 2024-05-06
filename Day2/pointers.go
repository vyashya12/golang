package main

import "fmt"

func main() {
	var s int = 10
	var t *int = &s

	fmt.Println("s=", s)
	fmt.Println("t=", t)
	fmt.Println("*t=", *t)
}

// *t prints the value at the memory address stored in t
