package main

import "fmt"

func main() {
	// integer
	var number int64 = 123123123123123123
	number2 := 123

	// error duplciate declare of variable
	// number2 := 321
	number2 = 321

	fmt.Println(number)

	number = 0

	fmt.Println(number)
	fmt.Println(number2)

	// float
	var float float32 = 123123.123
	float2 := 12312.12

	fmt.Println(float)
	fmt.Println(float2)

	// unsigned integer
	var unsigned uint8 = 123
	unsigned2 := 123123
	unsigned3 := uint64(12312312)

	fmt.Println(unsigned)
	fmt.Println(unsigned2)
	fmt.Println(unsigned3)

	// string
	var hello string = "Hello"
	hello = "world"
	hello2 := "Mantap"

	fmt.Println(hello)
	fmt.Println(hello2)

	// boolean
	var isTrue bool = false
	isTrueSecond := true

	fmt.Println(isTrue)
	fmt.Println(isTrueSecond)

	// constants
	const pi = 3.14
	const pi2 float32 = 3.14
	const pi3 float64 = 3.14

	fmt.Println(pi)
	fmt.Println(pi2)
	fmt.Println(pi3)

	const greating = "Hello"
	const greating2 = "World"

	fmt.Println(greating)
	fmt.Println(greating2)

	fmt.Printf("%s %s\n", greating, greating2) // string interpolation

	// multiple declare
	// var fourth, fifth, sixth string = "empat", "lima", "enam"
	fourth, fifth, sixth := "empat", "lima", "enam"

	fmt.Println(fourth)
	fmt.Println(fifth)
	fmt.Println(sixth)

	// underscore variable
	// used for unused variable
	_ = 123
	fmt.Println("Hello World")

	// declare variable with new
	// used for allocate memory (pointer)
	name := new(string)
	fmt.Println(name)
}
