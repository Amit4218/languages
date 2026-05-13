package main

import "fmt"

func main() {
	fmt.Println("________Arrays_________")

	var fruits [3]string
	fruits[0] = "apple"
	fruits[2] = "mango"

	fmt.Println("Fruits: ", fruits)
	fmt.Println("len of Fruits: ", len(fruits))

	var toys = [3]string{"car", "bike", "drum"}

	fmt.Println("Toys: ", toys)
	fmt.Println("len of Toys: ", len(toys))

}
