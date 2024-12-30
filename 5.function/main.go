package main

import "fmt"

func greet(name string) string {
	return "Hello " + name
}

func calculate(a, b int) (int, int) {
	return a + b, a * b
}

func main() {
	// function
	message := greet("Nabbil")
	fmt.Println(message)

	// multiple return values
	sum, product := calculate(2, 3)
	fmt.Println("Sum: ", sum, "Product: ", product)
}
