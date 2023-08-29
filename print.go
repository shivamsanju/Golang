package main

import "fmt"

func printer() {
	age := 23
	fmt.Printf("my age is %v \n", age)
	fmt.Printf("age is of type %T \n", age) // get type of variable
	fmt.Printf("you scored %f points \n", 20.23)
	fmt.Printf("you scored %0.2f points \n", 20.23) // rounded

	// save formatted strings
	name := fmt.Sprintf("my age is %v \n", age)
	fmt.Printf(name)
}
