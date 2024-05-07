package shapes

import "math"

func AreaOfCircle(radius float64) float64 {
	return math.Pi * radius * radius
}

func diameterOfCircle(radius float64) float64 {
	return 2 * radius
}

// Creating a package under a folder, so the package is shapes and the parent folder is shapes.
// We automatically export the functions within if they start with a capital letter
