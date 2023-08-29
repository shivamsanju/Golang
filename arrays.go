package main

import "fmt"

func arr() {
	// arrays -> fixed length
	var ages [3]int = [3]int{10, 20, 30}
	var ages2 = [3]int{10, 20, 30}
	names := [4]string{"yoshi", "mario", "peach", "bowser"}

	fmt.Println(ages, ages2, len(ages))
	fmt.Println(names, len(names), names[1])

	names[1] = "luigi"
	fmt.Println(names)

	// slices (like lists) (use arrays under the hood)
	var scores = []int{100, 200, 600}
	fmt.Println(scores, len(scores))
	scores[2] = 25
	fmt.Println(scores)

	scores = append(scores, 200)
	fmt.Println(scores, len(scores))

	// slice ranges
	rangeOne := names[1:3]
	rangeTwo := names[1:]
	rangeThree := names[:3]
	fmt.Println(rangeOne, rangeTwo, rangeThree)
}
