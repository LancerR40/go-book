package main

// import "fmt"

func main() {

	/*
		NOTE: You could optimize the memory by assigning types to variables with numeric types
	*/

	const intValue = 5 // it's assigned a type depending on of the OS architecture (32/64 bits)

	const int8Value int8 = 127   // 8 bits = -128 a 127
	const int16Value int16 = 350 // 16 bits = -2^15 a 2^15-1
	const int32Value int32 = 400 // 32 bits = -2^31 a 2^31-1
	const int64Value int64 = 500 // 64 bits = -2^63 a 2^63-1

	// For optimize more the memory, if you know the number is positive you should place a "uint" type

	const uintValue = 10 // it's assigned a type depending on of the OS architecture (32/64 bits)

	const uint8Value int8 = 127   // 8 bits = -128 a 127
	const uint16Value int16 = 350 // 16 bits = 0 a 2^15-1
	const uint32Value int32 = 400 // 32 bits = 0 a 2^31-1
	const uint64Value int64 = 500 // 64 bits = 0 a 2^63-1

	/* -------------------------------------- For the float values -------------------------------------- */

	const float32Value float32 = 20.00 // 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	const float64Value float64 = 40.00 // 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	/* -------------------------------------- For the string values -------------------------------------- */

	const stringValue string = "This is a string value in Go/Golang"

	/* -------------------------------------- For the boolean values -------------------------------------- */

	const itsFalse bool = false
	const itsTrue = true

	/* -------------------------------------- For the complex numbers values -------------------------------------- */

	const complexFloat32Value complex64 = 10 + 8i     // 32 bits = real and imaginary number of 32 bits
	const complexFloat64Value complex128 = 1000 + 10i // 64 bits = real and imaginary number of 64 bits

}
