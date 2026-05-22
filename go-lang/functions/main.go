package main

import "fmt"

func main() {
	fmt.Println("_________Functions________")
	fmt.Println("This itself is a function")
	greeter()
	result := add(2, 4)
	fmt.Println(result)
	rs := summ(2, 3, 4, 5, 6, 7, 8)
	fmt.Println(rs)
}

func add(a int, b int) int {
	return a + b
}

func summ(values ...int) int {
	total := 0
	for _, v := range values {
		total += v
	}
	return total
}

func greeter() {
	fmt.Println("Hello")
}
