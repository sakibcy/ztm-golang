//--Summary:
//  Create a program to calculate the area and perimeter
//  of a rectangle.
//
//--Requirements:
//* Create a rectangle structure containing a length and width field
//* Using functions, calculate the area and perimeter of a rectangle,
//  - Print the results to the terminal
//  - The functions must use the rectangle structure as the function parameter
//* After performing the above requirements, double the size
//  of the existing rectangle and repeat the calculations
//  - Print the new results to the terminal
//
//--Notes:
//* The area of a rectangle is length*width
//* The perimeter of a rectangle is the sum of the lengths of all sides

package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Gov          bool
}

type Bus struct {
	FrontSeat Passenger
}

type Coordinate struct {
	height int
	width  int
}

type Rectangle struct {
	a Coordinate
	b Coordinate
}

func area(rect Rectangle) int {
	return (rect.a.height * rect.a.width) + (rect.b.height * rect.b.width)
}

func main() {
	dora := Passenger{"Dora", 1, true}
	fmt.Println(dora)
	dora.Name = "HI"
	fmt.Println(dora)

	var (
		nobi = Passenger{Name: "Nobi", TicketNumber: 2} // if i don't pass other value then have to specify names
		sizu = Passenger{Name: "Sizu", TicketNumber: 4}
	)
	fmt.Println(nobi, sizu)

	var soneo Passenger
	soneo.Name = "Soneo"
	soneo.TicketNumber = 5

	// bus := Bus{FrontSeat: dora}
	bus := Bus{dora}
	fmt.Println(bus.FrontSeat.Name)

	react := Rectangle{a: Coordinate{1, 2}, b: Coordinate{4, 5}}

	fmt.Println("Area", area(react))
}
