package main

import "fmt"

func increment(x *int) {
	*x++
}

func main() {

	// efficiency for large structs or arrays with pointers
	// mutability: able to modify
	// passing pointers: pass pointer to function instead of the data so that it does not copy the data
	// type specific: even the pointer in go is type specific, pointer for string is different to pointer for int
	// using & operator: & operator passes the memory address of variable
	// using * operator: the * is used to get the value of variable at memory address provided by the pointer

	fmt.Println("Declare a pointer")
	var p *int
	fmt.Println(p)

	fmt.Println("Pointer to an existing variable")
	var x int = 10
	p1 := &x
	fmt.Println(p1)
	fmt.Println(*p1)

	fmt.Println("Create a pointer with new keyword")
	p2 := new(int)
	// we initialize the pointer
	*p2 = 100
	// the moment we say *p2 then it is storing in p2, memory address points to it and modifies i at same memory location
	fmt.Println(*p2)

	fmt.Println("Pointer to pointer")
	// nonsense, why would i have multiple levels of indirection
	var d int
	var p3 *int = &d
	var pp **int = &p3
	// also if there is a pointer pointing to another pointer, we need to use two asterisks
	fmt.Println(**pp)

	// common operations with pointers
	fmt.Println("Common Operations")
	var c int = 10
	var o *int = &c
	fmt.Println(o)

	// changing the value using a pointer
	// when we have a pointer and then we change the value using *, it changes the original variable value as well, because the pointer is showing the same memory location value to change
	*o = 20
	fmt.Println("o", o)
	fmt.Println("c=", c)

	// comparing pointers
	var a, b int
	t1 := &a
	t2 := &b
	t3 := &a

	// we can use the pointer to check if it is pointing to the same memory location, which in turn is the same value or not
	fmt.Println(t1 == t2)
	fmt.Println(t1 == t3)
	*t3 = 10
	fmt.Println(t1 == t3)

	fmt.Println("Passing pointer to functions")
	var number int = 10
	increment(&number)
	fmt.Println(number)
}
