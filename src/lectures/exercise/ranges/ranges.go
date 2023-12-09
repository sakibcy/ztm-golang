package main

import "fmt"

func main() {
	slice := []string{"Hello", "Wordl", "!"}

	// for i := 0; i < len(slice); i++ {
	// 	slice[0][i] // not posible in go
	// }

	for i, element := range slice {
		fmt.Println(i, element)

		for _, ch := range element {
			fmt.Printf("%q\n", ch)
		}
	}
}
