package main

import "fmt"

func main() {

	fmt.Println("_________Structs_________")
	// no inheritance in golang; No super or parent
	// Basically No Classes

	User := User{"Amit", "amit@mail.com", true, 67}
	// fmt.Println(User)

	fmt.Printf("User Details: %+v\n", User)
	fmt.Printf("User Name: %v\n", User.Name)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
