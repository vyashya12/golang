package main

import "fmt"

func main() {
	fmt.Println("Type Alias:")
	type ByteSize int64
	var b ByteSize = 1024

	fmt.Println(b)
}
