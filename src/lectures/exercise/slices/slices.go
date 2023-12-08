//--Summary:
//  Create a program to manage parts on an assembly line.
//
//--Requirements:
//* Using a slice, create an assembly line that contains type Part
//* Create a function to print out the contents of the assembly line
//* Perform the following:
//  - Create an assembly line having any three parts
//  - Add two new parts to the line
//  - Slice the assembly line so it contains only the two new parts
//  - Print out the contents of the assembly line at each step
//--Notes:
//* Your program output should list 3 parts, then 5 parts, then 2 parts

package main

import "fmt"

func printRoute(slice []string) {
	fmt.Println()
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	route := []string{"Home", "Grocery Store", "Department Store"}
	printRoute(route)

	route = append(route, "Home")
	printRoute(route)

	route = append(route, route...)
	printRoute(route)

	route = route[1:2]
	printRoute(route)
}
