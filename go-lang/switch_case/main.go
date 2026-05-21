package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("________Switch_CASE________")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("value of dice is: ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice Value is 1 and you can open")

	case 2:
		fmt.Println("Dice Value is 2 and you can move 2 spot")
		fallthrough
	case 3:
		fmt.Println("Dice Value is 3 and you can move 3 spot")
	case 4:
		fmt.Println("Dice Value is 4 and you can move 4 spot")
	case 5:
		fmt.Println("Dice Value is 5 and you can move 5 spot")
	case 6:
		fmt.Println("Dice Value is 6 and you can move 6 spot and roll again")
	}

}
