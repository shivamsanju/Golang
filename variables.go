package main

import "fmt"

func variablesLesson1() {
	// strings
	var nameOne string = "mario"
	var nameTwo = "luigi"
	var nameThree string

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "peach"
	nameThree = "bowser"
	nameFour := "yoshi" // fourth way to declare variable (can't be used outside a function)

	fmt.Println(nameOne, nameTwo, nameThree, nameFour)

	// ints
	var ageOne int = 20 //32 bits
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits and memory - int8 int16 int32 int64
	var numOne int8 = 17
	var numTwo int16 = -123
	var numThree uint8 = 231 // unsigned integer

	fmt.Println(numOne, numTwo, numThree)

	// float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 257285682567573725672678.98

	fmt.Println(scoreOne, scoreTwo)

}
