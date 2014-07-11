package main

import (
	"fmt"
)

func main() {
	sum_of_squares := 0
	square_of_sum := 0

	for i := 1; i <= 100; i++ {
		sum_of_squares += i * i
		square_of_sum += i
	}

	square_of_sum *= square_of_sum

	fmt.Println("Sum of squares:", sum_of_squares)
	fmt.Println("Square of sum:", square_of_sum)
	fmt.Println("Difference:", square_of_sum-sum_of_squares)
}
