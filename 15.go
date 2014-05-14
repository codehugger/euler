package main

import (
	"fmt"
)

func main() {
	fmt.Println("Problem 15")

	grid := [21][21]int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = 1

			lattice_sum := 0

			if j > 0 {
				lattice_sum += grid[i][j-1]
			}

			if i > 0 {
				lattice_sum += grid[i-1][j]
			}

			if i > 0 || j > 0 {
				grid[i][j] = lattice_sum
			}
		}
	}

	fmt.Println(grid[20][20])
}
