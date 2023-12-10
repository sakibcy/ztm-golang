package main

import "fmt"

// take lots of argument of same type
func sum(nums ...int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}

func main() {
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}

	b = append(b, a...)
	result := sum(b...)
	fmt.Println(result)

	ans := sum(1, 2, 3, 4)
	fmt.Println(ans)
}
