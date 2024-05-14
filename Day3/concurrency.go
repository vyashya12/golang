package main

import (
	"errors"
	"fmt"
)

func multiply(a int, b int) (int, error) {
	if a == 0 || b == 0 {
		return 0, errors.New("Answer will be zero")
	}
	return a * b, nil
}
func main() {
	// key features
	// goroutines and channels
	// goroutine is a lightweigth thread
	// more effiecient than traditional thread
	// non blocking
	// shared memory

	//create goroutine

	fmt.Println(multiply(4, 5))

}
