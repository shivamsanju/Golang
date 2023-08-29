package main

import (
	"fmt"
	"sort"
	"strings"
)

func useStdLib() {

	// STRINGS
	greeting := "Hello there friend!"

	fmt.Println(strings.Contains(greeting, "Hello"))

	fmt.Println(strings.ReplaceAll(greeting, "Hello", "Hi")) // doesn't change the original string

	fmt.Println(strings.ToUpper("lowercase string"))

	fmt.Println(strings.Index("lowercase string", "s")) // position of first s

	splitArr := strings.Split(greeting, " ")
	fmt.Println(splitArr, len(splitArr))

	// SORT
	ages := []int{10, 19, 28, 73, 15, 10, 12}
	sort.Ints(ages) // alters the original variable
	fmt.Println(ages)

	index := sort.SearchInts(ages, 10)
	index2 := sort.SearchInts(ages, 1022)
	fmt.Println(index, index2, len(ages))
}
