//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:
//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)
//* Create a map using the server names as the key and the server status
//  as the value
//* Set all of the server statuses to `Online` when creating the map
//* After creating the map, perform the following actions:
//  - call display server info function
//  - change server status of `darkstar` to `Retired`
//  - change server status of `aiur` to `Offline`
//  - call display server info function
//  - change server status of all servers to `Maintenance`
//  - call display server info function

package main

import (
	"fmt"
)

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

func printServerStatus(server map[string]int) {
	fmt.Println("There are", len(server), " servers")
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}
	// myMap := map[string]int{
	// 	"item 1": 18,
	// 	"item 2": 2,
	// 	"item 3": 3,
	// }

	myMap := make(map[string]int)

	myMap["A"] = 40
	myMap["B"] = 60
	myMap["C"] = 90

	fmt.Println(myMap["item 1"])
	fmt.Println(myMap["A"])
	fmt.Println(myMap["C"])

	delete(myMap, "C")
	fmt.Println(myMap["C"])

	val, found := myMap["B"]
	fmt.Println(val, found)
	fmt.Println()

	fmt.Println("Printing key value pairs")
	for key, value := range myMap {
		fmt.Println(key, value)
	}

	serverStatus := make(map[string]int)

	for _, server := range servers {
		serverStatus[server] = Online
	}

	fmt.Println()
	fmt.Println("Printing servers")
	printServerStatus(serverStatus)
}
