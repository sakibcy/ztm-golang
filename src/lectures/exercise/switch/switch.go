//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

import "fmt"

func main() {
	result := 1

	switch result {
	case 10, 1:
		fmt.Println("1")
		// fallthrough // will continue checking the next case
	case 2:
		fmt.Println("2")
	default:
		fmt.Println("20")
	}
}
