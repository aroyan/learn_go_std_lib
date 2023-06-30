package main

import "fmt"

type Person struct {
	name   string
	gender string
}

func main() {
	fruits := []string{"Apple", "Grape", "Banana"}

	customers := []Person{
		{name: "Winter", gender: "F"},
		{gender: "F", name: "Jurin"},
	}

	fmt.Printf("%v\n", fruits)

	// Print structs value only
	fmt.Printf("%v\n", customers)

	// Print structs with the fields name
	fmt.Printf("%+v\n", customers)

	// Print data types
	fmt.Printf("%T\n", customers)
	fmt.Printf("%T\n", fruits)
}
