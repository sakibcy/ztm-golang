//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Coordinate struct {
	X, Y int
}

// Regular function
func shiftBy(x, y int, coordinate *Coordinate) {
	coordinate.X += x
	coordinate.Y += y
}

// Pointer Receiver Function
// it create a option to get this function with '.' like OOP
func (coordinate *Coordinate) shiftBy(x, y int) {
	coordinate.X += x
	coordinate.Y += y
}

// Value Receiver Function
// it create a option to get this function with '.' like OOP
func (c Coordinate) shiftDist(other Coordinate) Coordinate {
	return Coordinate{other.X - c.X, other.Y - c.Y}
}

func main() {
	// =================== Pointer Receiver ===================== //
	coordinate := Coordinate{5, 5}

	// Regular function
	shiftBy(1, 1, &coordinate)
	fmt.Print("Regular -> X:", coordinate.X, " Y:", coordinate.Y)
	fmt.Println()

	// Pointer Receiver function
	coordinate.shiftBy(2, 2)
	fmt.Print("Pointer receiver -> X:", coordinate.X, " Y:", coordinate.Y)
	fmt.Println()

	// =========================== Value Reciever ===================== //
	first := Coordinate{6, 6}
	second := Coordinate{9, 9}

	distance := first.shiftDist(second)
	fmt.Println("Distance:", distance.X, distance.Y)
}
