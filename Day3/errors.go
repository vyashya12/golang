package main

import (
	"errors"
	"fmt"
)

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

// MyError is a custom error type
type MyError struct {
	Msg  string `json:"msg"`
	Code int    `json:"code"`
}

// Errr implements the errror interface for MyError
func (e *MyError) Error() string {
	return e.Msg
}

func SomeFunction(flag bool) error {
	if !flag {
		return &MyError{"Custom error occured", 500}
	}
	return &MyError{"No error", 200}
}

func SafeFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic", r)
		}
	}()
	panic("A problem occured")
}

func riskyFunction() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovere from: ", r)
		}
	}()
	panic("Something bad happened")
}

func main() {

	fmt.Println(divide(10, 0))
	fmt.Println(divide(10, 2))

	result, err := divide(50, 10)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	} else {
		fmt.Println("Result: ", result)
	}

	res := SomeFunction(true)
	if res != nil {
		fmt.Println(res)
	}

	SafeFunction()
}
