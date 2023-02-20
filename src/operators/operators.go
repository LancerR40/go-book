package main

import (
	"fmt"
	"math"
)

func main() {
	x := 10
	y := 50

	// Addition
	result := x + y

	fmt.Println("Addition", result)

	// Sustraction
	result = y - x

	fmt.Println("Subtraction", result)

	// Multiplication
	result = x * y

	fmt.Println("Multiplication", result)

	// Division
	result = y / x

	fmt.Println("Division", result)

	// Module
	result = y % x

	fmt.Println("Module", result)

	// Incremental
	x++

	fmt.Println("Incremental", x)

	// Decremental
	x--

	fmt.Println("Decremental", x)

	// - Tasks: Calcule the area of a rectangle, trapezoid and circle

	// Caculate the rectangle area
	var rectangleLength float32 = 40.5
	var rectangleWidth float32 = 50.2

	// Using := to declare variable, the type of its automatically defined
	rectangleArea := rectangleLength * rectangleWidth

	fmt.Println("Rectangle area", rectangleArea)

	// Caculate the trapezoid area
	trapezoidBase1 := 10.2
	trapezoidBase2 := 5.0
	trapezoidHeight := 20.0

	trapezoidArea := (trapezoidBase1 + trapezoidBase2) * (trapezoidHeight / 2)

	fmt.Println("Trapezoid area", trapezoidArea)

	// Calculate circle area
	circleRadio := math.Pow(10.2, 2)
	circleArea := math.Pi * circleRadio

	fmt.Println("Circle area", circleArea)
}
