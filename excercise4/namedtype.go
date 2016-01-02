// Exercise 4 - Named Type problem
//
// Declare a named type called counter with a base type of int.
// Declare a variable named c of type counter set to its zero value.
// Display the value of this variable and the variables type.
//
package main

import "fmt"

// counter is a named type for counting.
type counter int

// main is the entry point for the application.
func main() {
	// Declare a variable of type counter.
	var steps counter

	// Display the value of steps
	fmt.Println(steps)
	fmt.Printf("type for steps %T\n", steps)

}
