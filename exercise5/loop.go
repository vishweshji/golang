// Exercise 5 - Display even numbers up to 100
//
// Write a program that uses the continue keyword
// in a for loop to print only even numbers up to 100.
//
package main

import "fmt"

// main is the entry point for the application.
func main() {

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else {
			continue
		}
	}
}
