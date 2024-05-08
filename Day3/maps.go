package main

import "fmt"

func main() {
	fmt.Println("Maps")
	// Efficient: insert and delete easily with efficiency
	// Flexibility
	// Dynamic Sizing: no fixed size
	// Reference Type meaning: you can pass a map to a function to modifiy it

	fmt.Println("Map without initialization")
	var m map[string]int
	fmt.Println(m)

	fmt.Println("Map with initialization")
	f := map[string]int{
		"apple":  0,
		"banana": 1,
	}
	fmt.Println(f)

	//make is the most common way to create a map
	l := make(map[string]int)
	l["apple"] = 3
	l["banana"] = 5
	fmt.Println(l)

	s := make(map[string]string, 5)
	fmt.Println(s)

	//structs as map values
	type Product struct {
		Name  string
		Price float64
	}

	products := make(map[int]Product)
	products[1] = Product{Name: "Koko", Price: 10}
	products[2] = Product{Name: "Milk", Price: 8}
	fmt.Println(products)

	// nested maps
	company := make(map[string]map[string]int)
	company["Google"] = map[string]int{"location": 2, "salary": 100}
	fmt.Println(company)

	// common operations
	// len as usual
	fmt.Println(len(m))

	// add elemtents to map
	products[0] = Product{Name: "Koko", Price: 9.90}

	// retrieving elements from a map
	value := products[1]
	fmt.Println(value)

	// cecking existence f element in map
	product, ok := products[2]
	fmt.Println(product, ok)

	delete(products, 2)
	fmt.Println(products)

	for key, val := range f {
		fmt.Printf("Key: %s, Val: %d\n", key, val)
	}

}
