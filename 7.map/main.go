package main

import "fmt"

func main() {
	person := map[string]string{
		"name": "Nabbil",
		"city": "Jakarta",
	}

	fmt.Println(person["name"])
	fmt.Println(person["city"])
}
