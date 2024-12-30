package main

import "fmt"

func greet(name string) string {
	return "Hello " + name
}

func main() {
	// function
	message := greet("Nabbil")
	fmt.Println(message)
}
