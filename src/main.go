package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	// Constants: Values that cannot be change
	const pi1 float64 = 3.14 // Here we define de const type as "float64"
	const pi2 = 3.1415       // We define a const without type

	fmt.Println("Pi 1:", pi1, "Pi 2:", pi2)

	// Ways to declare variables
	base := 12          // Declared variable without type
	var height int = 14 // Declared variable with type
	var area int        // Declared variable without initial value, therefore the default value depend on of its type in this case it is "int"

	fmt.Println(base, height, area)

	// Zero values: Defined values without initial values, therefore it's definition depend on it's type
	var someNumber int      // It's default value is an "int"
	var someFloat64 float64 // It's default value is an "float 64"
	var someString string   // It's default value is a "string"
	var someBoolean bool    // It's default value is a "boolean (false by default)"

	fmt.Println(someNumber, someFloat64, someString, someBoolean)

	// Calculate the square area
	const squareBase = 10
	squareArea := squareBase * squareBase
	fmt.Println("Square area:", squareArea)
}
