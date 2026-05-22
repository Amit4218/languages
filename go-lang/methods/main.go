package main

import "fmt"

func main() {
	fmt.Println("_________Methods________")

	user := User{
		IsActive: true,
		Name:     "Amit",
		Email:    "amit@amit.com",
		Age:      89,
	}

	fmt.Printf("Original User: %v\n", user)

	user = user.UpdateUser("init")
	fmt.Printf("Updated User: %v\n", user)
}

type User struct {
	Name     string
	Age      int
	Email    string
	IsActive bool
}

func (u User) UpdateUser(name string) User {
	u.Name = name
	u.Email = name + "@" + name + ".com"
	return u
}
