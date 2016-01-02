// Exercise 2 - Simple problems on variables
//
// 1. Declare three variables that are initialized to their zero value.
// Declare a variable of type string, int and bool. Display the values of those variables.
//
// 2. Declare a new variable of type float32 and initialize the variable by
// converting the literal value of Pi (3.14).

package main

import "fmt"

// main is the entry point for the application.
func main() {
	// Declare variables that are set to their zero value.
	var id int
	var name string
	var isMinor bool

	// Display the value of those variables.
	fmt.Println(id)
	fmt.Println(name)
	fmt.Println(isMinor)

	// Declare variables and initialize.
	// Using the short variable declaration operator.
	month := 1
	day := "Saturday"
	working := false

	// Display the value of those variables.
	fmt.Println(month)
	fmt.Println(day)
	fmt.Println(working)

	// Perform a type conversion.
	pi := float32(3.14)

	// Display the value of that variable.
	fmt.Printf("%T [%v]\n", pi, pi)
}
