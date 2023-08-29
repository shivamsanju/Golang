package main

import "fmt"

func useLoops() {
	// while loop (we use for keyword only)
	x := 0
	for x < 5 {
		fmt.Println("value of x is: ", x)
		x++
	}

	for i := 0; i < 5; i++ {
		fmt.Println("value of i is: ", i)
	}

	names := []string{"mario", "luigi", "yoshi", "peach"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for index, val := range names {
		fmt.Println("names[", index, "] = ", val)
	}

	// if we don't wn't index
	for _, val := range names {
		fmt.Println(val)
	}

}
