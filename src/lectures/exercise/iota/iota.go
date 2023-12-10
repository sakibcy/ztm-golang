//--Summary:
//  Create a calculator that can perform basic mathematical operations.
//
//--Requirements:
//* Mathematical operations must be defined as constants using iota
//* Write a receiver function that performs the mathematical operation
//  on two operands
//* Operations required:
//  - Add, Subtract, Multiply, Divide
//* The existing function calls in main() represent the API and cannot be changed
//
//--Notes:
//* Your program is complete when it compiles and prints the correct results

package main

import "fmt"

const (
	Online = iota
	Offline
	Member
	Age
)

// starts at 3
const (
	Hello = iota + 3 // 3 = i
	i4               // 4
	i5               // 5
)

const (
	s1 = iota // 0
	_         // 1 (skip)
	_         // 2 (skip)
	s3        // 3
	s4        // 4
)

func main() {
	// iota is like enum in other languages
	fmt.Println(i4)

	// fmt.Println(add.calculate(2, 2)) // = 4

	// fmt.Println(sub.calculate(10, 3)) // = 7

	// fmt.Println(mul.calculate(3, 3)) // = 9

	// fmt.Println(div.calculate(100, 2)) // = 50
}
