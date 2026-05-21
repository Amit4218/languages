package main

import "fmt"

func main() {
	fmt.Println("________Loops_______")

	// days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	fmt.Printf("Day is %v \n", day)
	// }

	// rogueValue := 1
	// for rogueValue < 10 {
	// 	fmt.Println("Value is: ", rogueValue)
	// 	rogueValue++ // if this is not here an infinite loop will start
	// }

	rogueValue := 1
	for rogueValue < 10 {

		if rogueValue == 5 {
			rogueValue++
			continue
		}

		if rogueValue == 3 {
			goto comeHere
		}

		if rogueValue == 8 {
			break
		}

		fmt.Println("Value is: ", rogueValue)
		rogueValue++
	}

	// using goto
comeHere:
	fmt.Println("Hi, There")

}
