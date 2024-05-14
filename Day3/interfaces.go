package main

import (
	"fmt"
	"net/http"
)

// Interfaces are like OOP classes
// They have inheritance and polymorphism
// So the below are interfaces, can be empty, can be multiple methods, can be single methods, can embed methods
// Interfaces as a constraint, you can create a generic one that takes in a method

type Number interface {
	int64 | float64
}

func Sum[T Number](a, b T) T {
	return a + b
}

// Single method interface
type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return d.Name + " says Woof"
}

// Embedding interefaces
type Animal interface {
	Speaker
	Move() string
}

type Bird1 struct {
	Name string
}

func (b Bird1) Speak() string {
	return b.Name + " says Tweety"
}

func (b Bird1) Move() string {
	return b.Name + " flies"
}

// Empty Interface
func PrintAnything(v interface{}) {
	fmt.Println(v)
}

// Interface with multiple methods
type Vehicle interface {
	Start() string
	Stop() string
}

type Car struct {
	Model string
}

func (c Car) Start() string {
	return c.Model + " car started"
}

func (c Car) Stop() string {
	return c.Model + " car stopped"
}

// Interface as a Constant
// func Describe[T Speaker](t T) {
// 	fmt.Println(t.Speak())
// }

// ----------------------------------------------------------------

type Human interface {
	Walks() string
	HasMoney() float64
	IsLegal() bool
}

type Person struct {
	Name  string
	Money float64
}

func (p Person) Walks() string {
	return p.Name + " walks"
}

func (p Person) HasMoney() float64 {
	return p.Money
}

func (p Person) IsLegal() bool {
	return true
}

// type assertions
func PrintDetails(i interface{}) {
	str, ok := i.(string)
	if ok {
		fmt.Println("Its string", str)
	} else {
		fmt.Println("Not a string")
	}
}

func Describe(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Println("Integer", v)
	case string:
		fmt.Println("String", v)
	default:
		fmt.Println("Unknown type")
	}
}

// ----------------------------------------------------------------
func main() {
	fmt.Println("Interfaces")
	// type agnostic : not tied to any specific type
	// satisfy the interface by implementing all the methods in interface

	// implicit implementation:  does not require any type to explicitly declare that it implements an interface

	// runtime polymorphisms: single interface can represent different types

	// decoupling : iterfaces helps in decoupling by not depending on the concrete implementation

	// empty interface: can hld any values of an type

	fmt.Println("Single method interfaces")
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	fmt.Println("Embedding interfaces")
	type Writer interface {
		Writer(p []byte) (n int, err error)
	}

	type ReadWriter interface {
		Reader
		Writer
	}

	fmt.Println("Empty Interface")
	type Any interface{}

	fmt.Println("Interface with multiple methods")
	type Shape interface {
		Area() float64
		Perimeter() float64
	}

	fmt.Println("Functional interface")
	// often used in Go for callback
	// only one method
	type HandlerFunc interface {
		ServeHTTP(w http.ResponseWriter, r *http.Request)
	}

	fmt.Println("Interface as Constraint")

	// above for understanding only, below are the actual methods and interfaces usage

	dog := Dog{Name: "Tommy"}
	bird := Bird1{Name: "Adam"}
	car := Car{Model: "Toyota"}

	fmt.Println(dog.Speak())
	fmt.Println(bird.Speak())
	fmt.Println(bird.Move())

	PrintAnything("An empty interface can hold anything")
	PrintAnything(43)

	fmt.Println(car.Start())
	fmt.Println(car.Stop())

	Describe(dog)
	Describe(bird)

	// common operations

	// ok so we create an interface with methods then we have a struct and then we have funcs that take in these structs and return these methods and we need to define what the method does in here

	// example on line 80
	yash := new(Person)
	yash.Name = "Yash"
	yash.Money = 3459.02
	var t Human = Person{Name: "Tom", Money: 200.44}
	r, ok := t.(Person)
	if ok {
		fmt.Println(r.HasMoney())
	}

	// Type assertions
	Describe("Hello World")
	Describe(2024)
	Describe(42)

	fmt.Println(yash.Walks())
	fmt.Println(yash.HasMoney())
	fmt.Println(yash.IsLegal())

}
