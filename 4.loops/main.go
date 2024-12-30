package main

import "fmt"

func main() {
	// for
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// while
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// infinite loop
	for {
		fmt.Println("Hello World")
		break // prevent infinite loop
	}
}
