package main

import "fmt"

func main() {
	fmt.Println("__________Defer________")

	// defer is executed in an lifo order
	// whenever the function or program is about to exit all the
	// everything in the defer stack is executed

	fmt.Println("Hello")

	defer fmt.Println("Defer executed")
	fmt.Println("First defer was before this line, call myloop now")
	myloop()
	fmt.Println("Loop has ended")
	fmt.Println("End of program")
}

func myloop() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}
