package main

import "fmt"

func main() {
	x := 2

	// if else
	if x > 3 {
		fmt.Println("x is greater than 3")
	} else {
		fmt.Println("x is not greater than 3")
	}

	// switch
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("x is not 1 or 2")
	}

	day := "saturday"

	switch day {
	case "monday":
		fmt.Println("You have to work")
	case "friday":
		fmt.Println("Prepare! you have to go home")
	case "saturday":
		fmt.Println("It's time learn golang")
	}
}
