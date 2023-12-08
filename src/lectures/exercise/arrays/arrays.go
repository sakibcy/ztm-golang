//--Summary:
//  Create a program that can store a shopping list and print
//  out information about the list.
//
//--Requirements:
//* Using an array, create a shopping list with enough room
//  for 4 products
//  - Products must include the price and the name
//* Insert 3 products into the array
//* Print to the terminal:
//  - The last item on the list
//  - The total number of items
//  - The total cost of the items
//* Add a fourth product to the list and print out the
//  information again

package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func isCleaned(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is celaned")
		} else {
			fmt.Println(room.name, "is not cleaned")
		}
	}
}

func main() {
	rooms := [...]Room{
		{name: "waiting"},
		{name: "office"},
		{name: "reception"},
		{name: "reception"},
	}

	isCleaned(rooms)

	rooms[0].cleaned = true
	isCleaned(rooms)
}
