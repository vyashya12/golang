package main

import "fmt"

// Employee struct is an employee record
type Employee struct {
	Name     string
	Position string
	Salary   float64
}

//Update salary
func UpdateSalary(emp *Employee, newSalary float64) {
	emp.Salary = newSalary
}

func main() {
	emp := Employee{
		Name:     "John",
		Position: "Software Engineer",
		Salary:   100000,
	}

	fmt.Println("Before", emp)

	UpdateSalary(&emp, 200000)

	fmt.Println("After", emp)
}
