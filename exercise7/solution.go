// Exercise 7+8 - Fahrenheit into Celsius using conversion Package
//
// Exercise 7 - Write a program that converts from Fahrenheit into Celsius. (C = (F - 32) * 5/9)
//
// Exercise 8 - Create a package called ftoc for the problem in Exercise 7.
// Write another Go program which uses the ftoc package.

package main

import (
	"fmt"
	"github.com/golang/conversion"
)

func main() {
	fmt.Println(conversion.Convert_FtoC(95))
}
