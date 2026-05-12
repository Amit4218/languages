package main

import "fmt"

func main() {
	var name string = "Amit"
	var number = 18
	no_type_decleration := "no_type_decleration" // can be any things like and int, string, bool etc

	fmt.Printf("%s -> This was declared using both var and type: \n", name)
	fmt.Printf("%d -> This was declared using only var: \n", number)
	fmt.Printf("%s -> This was declared using only walrus operator ':=' no var or type: \n", no_type_decleration)
}
