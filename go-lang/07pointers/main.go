package main

import "fmt"

func main() {
	fmt.Println("____________Pointers_________")

	var ptr *int
	fmt.Println("Value of pointer is ", ptr)

	myNumber := 2

	// refrencing to mynumber
	var myPtr = &myNumber
	fmt.Println("Value of pointer is ", myPtr)
	fmt.Println("Value of pointer is using * ", *myPtr)

	// performing operation
	*myPtr = *myPtr * 2
	fmt.Println("New value of the myNumber is: ", myNumber)
}
