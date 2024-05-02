package main

import "fmt"

func main() {
	// 5 ways to declare variables
	// Using var
	var name string
	var age int
	var z bool

	// Declare multiple variables
	var (
		firstname, lastName string
		year                int
	)

	// With values
	var sentence string = "Hello there"
	var number int = 10

	// Type inference
	var meesage = "New message"

	// Short variable declaration in function
	note := "Hello there note"
	odd_number := 20

	const Pi = 3.14

	fmt.Println(name, age, z, firstname, lastName, year, sentence, number, meesage, note, odd_number, Pi)
}

// Zero values
// name = empty
// age = 0
// z = false
// firstname = empty
// lastName = empty
// sentence = "Hello there"
// const cannot be reassigned
