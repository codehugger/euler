package main

import (
	"fmt"
	"math"
)

func triangle_sum(n int) int {
	return n * (n + 1) / 2
}

func reverse_triangle_sum(n float64) float64 {
	return 0.5*math.Sqrt(8*n+1) - 0.5
}

func divisors(x int) int {
	limit := x
	numberOfDivisors := 0

	for i := 1; i < limit; i++ {
		if x%i == 0 {
			limit = x / i
			if limit != i {
				numberOfDivisors++
			}
			numberOfDivisors++
		}
	}
	return numberOfDivisors
}

func factors(n int) []int {
	var f []int
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			f = append(f, i)
		}
	}
	f = append(f, n)
	return f
}

func main() {
	fmt.Println("Problem 12")

	number := 1
	triangle_factors := 0

	for triangle_factors < 500 {
		sum := triangle_sum(number)
		num_factors := divisors(sum)
		if num_factors > triangle_factors {
			triangle_factors = num_factors
		}
		number += 1
	}

	fmt.Println(number)
}
