package main

import "fmt"

func useConditionals() {
	age := 45

	fmt.Println(age < 50)
	fmt.Println(age > 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("Age is less than 30")
	} else if age < 40 {
		fmt.Println("Age is less than 40")
	} else {
		fmt.Println("Age is greater than 40")
	}
}
