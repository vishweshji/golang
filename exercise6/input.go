// Exercise 6 - Take User Input
//
// Write a program that uses the Scanf function from the fmt package to read user input.
// This program should take in a number entered by the user and double it. Display the result.
//
package main

import "fmt"

// main is the entry point for the application.
func main() {

	// Accept a number from user
	var age int
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("In another ", age, " years your age will be :", age*2)
}
