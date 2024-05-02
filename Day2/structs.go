package main

import "fmt"

func main() {
	fmt.Println("Hello Structs")
	type Person struct {
		Name string
		Age  int
	}

	var q Person = Person{"Yash", 23}
	fmt.Println("q=", q)
	fmt.Println(q.Name)
	fmt.Println(q.Age)
}
