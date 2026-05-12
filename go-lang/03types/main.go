package main

import "fmt"

func main() {

	var integer int = 99
	var str string = "This is a string!"
	var isTrue bool = true
	var floating float32 = 69.99

	fmt.Printf("%d -> There are multiple type of int like u8, u16, u32, u64 but mostly int is used\n", integer)
	fmt.Printf("%s -> Anything inside single brackets or double are strings \n", str)
	fmt.Printf("%.2f -> There are two types of float point number float32 and float64 \n", floating)
	fmt.Printf("%v -> booleans are either True or False \n", isTrue)
}
