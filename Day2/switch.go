package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Basic Switch")
	day := 3
	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid Day")
	}

	fmt.Println("Switch with initialization")
	switch today := time.Now().Weekday(); today {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	fmt.Println("Switch without a condition (liek if-else)")
	x := 42
	switch {
	case x < 0:
		fmt.Println("x is Negative")
	case x == 0:
		fmt.Println("x is Zero")
	default:
		fmt.Println("x is Positive")
	}
}
