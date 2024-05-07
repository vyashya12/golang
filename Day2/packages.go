package main

import (
	"day2/shapes"
	"fmt"
)

func main() {
	radius := 5.0
	areaCircle := shapes.AreaOfCircle(radius)
	fmt.Printf("Area of Circle: %v\n", areaCircle)

	side := 4.0
	araeSquare := shapes.AreaOfSquare(side)
	fmt.Printf("Area of Square: %v\n", araeSquare)
}

// to use the functions from another package. for example here we are using package shapes, and importing it.
// this is the way we need to import it
