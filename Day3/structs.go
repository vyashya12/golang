package main

import (
	"fmt"
	"reflect"
)

type Bird struct {
	Species string `json:"species"`
}

func (a *Bird) Greet() string {
	return "Hello " + a.Species
}

func main() {
	fmt.Println("Structs")
	// structs are used to model real world data structures

	// Also often used to hold configuration data

	// common design patterns like Decorator, Adapter, etc can be implemented using structs

	// fields make up a struct, a fied can be any type

	// we can define methods on structs

	// tags : for metadata purposes

	fmt.Println("Basic Struct Declarations")
	type Person struct {
		Name string
		Age  int
	}

	fmt.Println("Declaring and initializing ")
	var product Person = Person{
		Name: "Yash",
		Age:  23,
	}
	fmt.Println(product)

	fmt.Println("Using the new Keyword")
	person2 := new(Person)
	person2.Name = "CK"
	person2.Age = 30

	// declare and one time use
	fmt.Println("Anonymous Structs")
	var personD = struct {
		Name string
		Age  int
	}{
		Name: "Daryl",
		Age:  29,
	}

	fmt.Println(personD)

	fmt.Println("Nested Structs")
	type Address struct {
		City, State string
	}

	type PersonLocate struct {
		Name    string
		Age     int
		Address Address
	}
	personLocate := PersonLocate{
		Name: "Happy",
		Age:  40,
		Address: Address{
			City:  "Bangalore",
			State: "Karnataka",
		},
	}

	fmt.Println(personLocate)

	fmt.Println("Embedded (Anonymous) Structs")
	type Contact struct {
		Phone, Email string
	}

	type PersonContact struct {
		Person
		Contact string // embedded field
	}

	personMobile := PersonContact{
		Person: Person{
			Name: "Shin",
			Age:  30,
		},
		Contact: "011-123123123",
	}

	fmt.Println("Embedded field stuct", personMobile)

	// common operations
	// Creating an instance of Struct
	type Animal struct {
		Species string
	}

	a := Animal{Species: "Bird"}
	fmt.Println(a)

	// accessing and modiying fields
	fmt.Println(a.Species)
	a.Species = "Mammal"
	fmt.Println(a)

	// Pointers to structs
	a2 := &Animal{Species: "Donkey"}
	fmt.Println(a2)

	// Methods on Structs
	// func (a Animal) Greet() string {
	// 	return "Hello " + a.Species
	// }

	// taggin on struct fields
	// type Bird struct {
	// 	Species string `json:"species"`
	// }

	b := Bird{Species: "Parrot"}
	fmt.Println(b)

	// calling the method that is attached to the struct
	fmt.Println(b.Greet())

	// getting the tag tht we added into the struct
	t := reflect.TypeOf(b)
	field, _ := t.FieldByName("Species")
	fmt.Println(field.Tag.Get("json"))

}
