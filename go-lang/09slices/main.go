package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("________Slices________")

	var fruits = []string{"apple", "banana"}

	fmt.Printf("Type of Fruits variable %T\n", fruits)
	fmt.Println("Fruits", fruits)

	fruits = append(fruits, "peach", "mango")
	fmt.Println("Fruits after append", fruits)

	fmt.Println("Fruits slices", fruits[1:3])

	fmt.Println("________Using Make________")

	var scores = make([]int, 2)

	scores[0] = 10
	scores[1] = 20
	// scores[3] = 30 // this will crash the code

	// this will reallocate the memory and make space automatically
	scores = append(scores, 3439, 982, 300, 1256, 21, 3411)
	fmt.Println(scores)

	fmt.Println("________Using Sorted Method________")
	fmt.Println(sort.IntsAreSorted(scores))
	sort.Ints(scores)
	fmt.Println(scores)

	fmt.Println("________Remove Value from slices________")

	var courses = []string{"reactjs", "javaScript", "swift", "python", "ruby"}
	fmt.Println(courses)

	var slice_index int = 2
	courses = append(courses[:slice_index], courses[slice_index+1:]...)
	fmt.Println(courses)

}
