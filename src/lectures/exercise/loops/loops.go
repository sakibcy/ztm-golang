//--Summary:
//  Implement the classic "FizzBuzz" problem using a `for` loop.
//
//--Requirements:
//* Print integers 1 to 50, except:
//  - Print "Fizz" if the integer is divisible by 3
//  - Print "Buzz" if the integer is divisible by 5
//  - Print "FizzBuzz" if the integer is divisible by both 3 and 5
//
//--Notes:
//* The remainder operator (%) can be used to determine divisibility

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// while loop
	num := 10
	for num > 0 {
		fmt.Println(num)
		num--
	}

	// Infinite loop
	var x int = 2
	for {
		if x == 3 {
			fmt.Println("X is now 3")
			break
		}
		fmt.Println("X is not 3")
		x++
	}

	// continue
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
}
