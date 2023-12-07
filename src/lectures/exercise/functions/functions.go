//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func double(x int) int {
	x *= 2
	return x
}

func greet() {
	fmt.Println("Hello user :)")
}

func greetPerson(name string) {
	fmt.Println("Hello", name)
}

func messageReturn() string {
	return "Message Return"
}

func addThreeNumber(x, y, z int) int {
	return x + y + z
}

func returnAnyNumber() int {
	return 4
}

func returnAnyTwoNumber() (int, int) {
	return 1, 2
}

func main() {
	greet()

	addition := add(1, 2)
	fmt.Println("Add tow nums:", addition)

	dozen := double(6)
	fmt.Println("The dozen is:", dozen)

	//--Requirements:
	//* Write a function that accepts a person's name as a function
	//  parameter and displays a greeting to that person.
	greetPerson("Dora")

	//* Write a function that returns any message, and call it from within
	//  fmt.Println()
	fmt.Println(messageReturn())

	//* Write a function to add 3 numbers together, supplied as arguments, and
	//  return the answer
	fmt.Println("Add three number:", addThreeNumber(1, 2, 3))

	//* Write a function that returns any number
	fmt.Println("Return any number", returnAnyNumber())

	//* Write a function that returns any two numbers
	a, b := returnAnyTwoNumber()
	fmt.Println("Return any two number", a, b)

	//* Add three numbers together using any combination of the existing functions.
	//  * Print the result
	//* Call every function at least once
}
