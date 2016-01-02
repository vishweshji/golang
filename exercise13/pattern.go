// Exercise 13 - Loop Pattern
//
// Create a Go program that prints the following (up to 10 characters):
// O
// OO
// OOO
// OOOO
// OOOOO
// OOOOOO
// OOOOOOO
// OOOOOOOO
// OOOOOOOOO
// OOOOOOOOOO
//

package main

import (
	"fmt"
	"strings"
)

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%s\n", strings.Repeat("0", i))
	}
}
