package main

import "fmt"

func main() {
	fmt.Println("Function as variable")
	var v func(int) int
	v = func(x int) int { return x * x }
	result := v(5)

	fmt.Println("result=", result)
}

//function as a variable, it was defined and then set a funtion inside then use it as line 9
