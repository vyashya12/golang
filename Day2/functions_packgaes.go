package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func sum_of_num(nums ...int) int {
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
}

func main() {
	result := add(3, 4)
	fmt.Println(result)

	a, b := swap("world", "hello")
	fmt.Println("a, b=", a, b)

	fmt.Println("Named Return Value")
	sum := 100
	x, y := split(sum)
	fmt.Printf("The split of %d is: x = %d, y = %d\n", sum, x, y)

	fmt.Println("Variadic functions")
	summed := sum_of_num(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11)
	fmt.Printf("The sum of nums functions returns %d", summed)

	fmt.Println("Hi world!!!!")
}
