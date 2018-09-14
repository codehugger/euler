package main

import (
	"fmt"
)

func main() {
	const WIDTH = 1001
	const HEIGHT = 1001
	circles := WIDTH / 2

	counter := 1
	posX := WIDTH / 2
	posY := HEIGHT / 2

	var spiral [WIDTH][HEIGHT]int

	// initial step of one
	spiral[WIDTH/2][HEIGHT/2] = counter

	for radius := 1; radius <= circles; radius++ {
		// go right until end
		for posX < WIDTH/2+radius {
			spiral[posY][posX] = counter
			posX++
			counter++
		}

		// go down until end
		for posY <= HEIGHT/2+radius {
			spiral[posY][posX] = counter
			posY++
			counter++
		}

		// move the pointer to the next starting point
		posX--
		posY--

		// // go left until end
		for posX >= WIDTH/2-radius {
			spiral[posY][posX] = counter
			counter++
			posX--
		}

		// move the pointer to the next starting point
		posY--
		posX++

		// go up until end
		for posY >= HEIGHT/2-radius {
			spiral[posY][posX] = counter
			posY--
			counter++
		}

		// move the pointer to the next starting point
		posY++
		posX++

		// go right until end
		for posX <= HEIGHT/2+radius {
			spiral[posY][posX] = counter
			posX++
			counter++
		}
	}

	diagonalSum := 0

	// diagonal down from left
	x := 0
	y := 0
	for x < WIDTH && y < HEIGHT {
		fmt.Println(spiral[y][x])
		diagonalSum += spiral[y][x]
		x++
		y++
	}

	// diagonal sum down from right
	x = WIDTH - 1
	y = 0
	for x >= 0 && y < HEIGHT {
		fmt.Println(spiral[y][x])
		diagonalSum += spiral[y][x]
		x--
		y++
	}

	// subtract one since we took the center twice
	diagonalSum -= 1

	fmt.Println("Problem 28")

	// for i, _ := range spiral {
	// 	fmt.Println(spiral[i])
	// }
	fmt.Printf("Diagonal Sum is %d\n", diagonalSum)
}
