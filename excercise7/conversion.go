// EExercise 7 - Fahrenheit into Celsius
//
// Write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
//

package main

import "fmt"

func main() {
	fahrenheit := 95
	celsius := (fahrenheit - 32) * 5 / 9
	fmt.Println(celsius)
}
