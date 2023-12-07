//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

import "fmt"

func main() {
	var myName = "Dora"
	fmt.Println("my name is: ", myName)

	var name string = "Nobi"
	fmt.Println("name = ", name)

	userName := "Doraemon" // create and assign
	fmt.Println("username: ", userName)

	var sum int // default value is 0
	fmt.Println(sum)

	part1, part2 := 3, 4
	fmt.Println("part 1 is:", part1, "part 2 is:", part2)

	// ok, error
	part3, part2 := 6, 5
	fmt.Println(part3, part2)

	var (
		id    = 30
		email = "e@e.com"
	)

	fmt.Println("id:", id, "email:", email)

	//* Store your favorite favoriteColor in a variable using the `var` keyword
	var favoriteColor = "red"
	fmt.Println("My favorite color is", favoriteColor)

	//* Store your birth year and age (in years) in two variables using
	//  compound assignment
	birthYear, ageInYears := 1990, 43
	fmt.Println("BirthYear:", birthYear, ",Age in Year:", ageInYears)

	//* Store your firstInitial & lastInitial initials in two variables using block assignment
	var (
		firstInitial string = "D"
		lastInitial  string = "N"
	)
	fmt.Println("Initials: ", firstInitial, lastInitial)

	//* Declare (but don't assign!) a variable for your ageInDays (in days),
	//  then assign it on the next line by multiplying 365 with the ageInDays
	// 	variable created earlier
	var ageInDays int
	ageInDays = ageInYears * 365
	fmt.Println("Age in Days:", ageInDays)
}
