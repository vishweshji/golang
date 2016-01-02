// Exercise 3 - Two problems on Constants
//
// 1. Declare an untyped and typed constant and display their values.
//
// 2.Multiply two literal constants into a typed variable and display
// the values of both the constants and variables.

package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare an untyped and typed constant.
	const hello = "Hello"
	const typedHello string = "世界"

	// Display the value of those variables.
	fmt.Println(hello)
	fmt.Println(typedHello)

	// Multiply two literal constants into a typed variable
	// Calcualte how many seconds in an hour (60 minutes)
	seconds := 60 * 60
	// Display the value of those variables.
	fmt.Println(seconds)
}
