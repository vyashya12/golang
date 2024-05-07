package main

import "fmt"

//characteristics of Go Arrays
// fixed size
// value type
// zero value

// passing arrays to functions
func processArray(arr10 [5]int) [5]int {
	fmt.Println(arr10)
	return arr10
}

func main() {
	//declaration without initialization
	// declaration is on the left side of the equals
	var arr [5]int
	fmt.Println("arr=", arr)

	// declaration with initializationq
	// initializationq is on the right side of the equals
	var arr2 [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2=", arr2)

	//short hand declaration
	// short hand does not need declaration type, can just use the initialization and the type will be set
	arr3 := [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr3=", arr3)

	// intialization with special element
	// sparse array
	// below left side of colon is the position of the value
	arr4 := [5]int{1: 10, 3: 30}
	fmt.Println("arr4", arr4)

	// no need to set size, it will infer from usage
	arr5 := [...]int{1, 2, 3, 4}
	fmt.Println("arr5", arr5)

	//multi dimensional arrays
	var multiArr [2][3]int
	fmt.Println("multiArr", multiArr)

	//multi dimensional arrays with initialization
	multiArr2 := [2][3]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println("multiArr2", multiArr2)

	//Arrays with pointers
	var arr6 [5]*int
	num1, num2, num3, num4, num5 := 13, 14, 15, 16, 17
	arr6[0] = &num1
	arr6[1] = &num2
	arr6[2] = &num3
	arr6[3] = &num4
	arr6[4] = &num5
	fmt.Println("arr6", arr6)
	for i := 0; i < len(arr6); i++ {
		fmt.Printf("pointerArray[%d] = %d\n", i, *arr6[i])
	}

	//array of structs
	type Person struct {
		Name string
		Age  int
	}

	var people [2]Person
	people[0] = Person{"Yash First", 23}
	people[1] = Person{"Yash Second", 23}
	fmt.Println("people", people[0])

	//common operations of arrays (methods)
	//length
	var myArr [5]int
	length := len(myArr)
	fmt.Println(length, myArr)

	//copy an array
	newArr1 := [3]int{1, 2, 3}
	var copiedArr = newArr1
	fmt.Println(newArr1, copiedArr)

	//modifying an array
	myArr[0] = 100
	fmt.Println(myArr)

	processArray(myArr)

	//comparing arrays
	isEqual := myArr == arr3
	fmt.Println(isEqual)
}
