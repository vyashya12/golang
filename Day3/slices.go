package main

import (
	"fmt"
	"sort"
)

func modifySlice(s []int) {
	if len(s) > 0 {
		s[0] = 999
	}
	fmt.Println("Inside modifySlice", s)
}

func main() {
	fmt.Println("Slices without initializing")
	var mySlice []int
	fmt.Println(mySlice)

	fmt.Println("Slices with shorthand and initializing")
	mySlice2 := []int{1, 2, 3}
	fmt.Println(mySlice2)

	fmt.Println("Slices with make and initializing")
	mySlice3 := make([]int, 5)
	fmt.Println(mySlice3)

	fmt.Println("Slices with make and initializing using the make function")
	mySlice4 := make([]int, 5, 10)
	fmt.Println(mySlice4)

	fmt.Println("Slice from array")
	// you can create slices from arrays
	// python array
	arr := [5]int{1, 2, 3, 4, 5}
	mySlicesFromArr := arr[1:4]
	fmt.Println(mySlicesFromArr)

	//Common operations for slice
	var slice1 = []int{1, 2, 3, 4, 5}
	length := len(slice1)
	fmt.Println(length)
	capacity := cap(slice1)
	fmt.Println(capacity)

	var numbers []int
	//Appending a single element
	numbers = append(numbers, 1)
	fmt.Println(numbers)
	//Appending multiple elements
	numbers = append(numbers, 2, 3, 4, 5, 1, 2, 6, 7, 8, 9, 12, 11)
	fmt.Println(numbers)

	//slicing a slice
	subSlice := numbers[1:4]
	firstTwo := numbers[:2]
	fromThird := numbers[2:]
	fmt.Println(subSlice)
	fmt.Println(firstTwo)
	fmt.Println(fromThird)

	// delete an element from slice
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)

	// checking if slice is empty
	if len(numbers) == 0 {
		fmt.Println("Slice is empty")
	}

	// sorting a slice
	fmt.Println("Slice sorting")
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] < numbers[j]
	})
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i] > numbers[j]
	})
	fmt.Println(numbers)

	//finding an element in a slice
	for index, v := range numbers {
		if v == 2 {
			fmt.Println("Element found")
			fmt.Println(index)
			break
		}
	}

	fmt.Println("Iterating over slices")
	iteSlice := []int{10, 20, 30, 40, 50, 60}
	for i := 0; i < len(iteSlice); i++ {
		fmt.Println(iteSlice[i])
	}
	fmt.Println("Iterating over slices idiomatic way")
	sort.Slice(iteSlice, func(i, j int) bool {
		return iteSlice[i] > iteSlice[j]
	})
	for _, value := range iteSlice {
		fmt.Println(value)
	}

	originalSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(originalSlice)

	modifySlice(originalSlice)

	fmt.Println(originalSlice)
}

// dynamic size
// reference type - any modifications to a slice passed into a function will be visible in the outside slice
