// Exercise 12 - Loop
//
// a. Create a simple loop with the for construct.
// Make it loop 5 times and print out the loop counter with the fmt package.
//
// b. Rewrite the above to use goto. for is not to be used here.
//
package main

import "fmt"

// main is the entry point for the application.
func main() {

	for i := 1; i <= 5; i++ {
		fmt.Println("Loop counter using for : ", i)
	}

	var counter int = 1
P:
	fmt.Println("Loop counter using goto : ", counter)
	counter++
	if counter <= 5 {
		goto P
	}

}
