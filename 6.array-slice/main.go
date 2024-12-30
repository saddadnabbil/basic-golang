package main

import "fmt"

func main() {
	// array
	var numbers [3]int = [3]int{10, 20, 30}
	fmt.Println(numbers)

	// slice
	numbersNew := []int{10, 20, 30}
	numbersNew = append(numbersNew, 40)
	fmt.Println(numbersNew)
}
