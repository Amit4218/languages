package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welome := "Learnign to take user Input"
	fmt.Println(welome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter something...")

	// comma ok || err ok syntax
	input, _ := reader.ReadString('\n')

	// if err != nil {
	// 	fmt.Println("An error occured")
	// }

	fmt.Println("You typed this", input)
	fmt.Printf("The input type %T", input)

}
