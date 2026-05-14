package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("__________CALCULATOR_________")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your first number: ")
	firstInput, _ := reader.ReadString('\n')

	fmt.Print("Enter your second number: ")
	secondInput, _ := reader.ReadString('\n')

	fmt.Print("Select operation (1:add 2:subtract 3:multiply 4:divide): ")
	operationInput, _ := reader.ReadString('\n')

	// Convert inputs
	firstNumber, err := strconv.ParseFloat(strings.TrimSpace(firstInput), 64)
	if err != nil {
		fmt.Println("Invalid first number")
		return
	}

	secondNumber, err := strconv.ParseFloat(strings.TrimSpace(secondInput), 64)
	if err != nil {
		fmt.Println("Invalid second number")
		return
	}

	operation, err := strconv.Atoi(strings.TrimSpace(operationInput))
	if err != nil {
		fmt.Println("Invalid operation")
		return
	}

	// Perform operation
	switch operation {
	case 1:
		fmt.Printf("Result: %.2f\n", firstNumber+secondNumber)

	case 2:
		fmt.Printf("Result: %.2f\n", firstNumber-secondNumber)

	case 3:
		fmt.Printf("Result: %.2f\n", firstNumber*secondNumber)

	case 4:
		if secondNumber == 0 {
			fmt.Println("Cannot divide by zero")
			return
		}
		fmt.Printf("Result: %.2f\n", firstNumber/secondNumber)

	default:
		fmt.Println("Invalid operation selected")
	}
}
