package main

import (
	"math/big"
)

/* Consider all integer combinations of ab for 2 ≤ a ≤ 5 and 2 ≤ b ≤ 5:
 *
 * 22^2=4, 23^2=8, 24^2=16, 25^2=32
 * 32^2=9, 33^2=27, 34^2=81, 35^2=243
 * 42^2=16, 43^2=64, 44^2=256, 45^2=1024
 * 52^2=25, 53^2=125, 54^2=625, 55^2=3125
 *
 * If they are then placed in numerical order, with any repeats removed,
 * we get the following sequence of 15 distinct terms:
 *
 * 4, 8, 9, 16, 25, 27, 32, 64, 81, 125, 243, 256, 625, 1024, 3125
 *
 * How many distinct terms are in the sequence generated by a^b for 2 ≤ a ≤ 100 and 2 ≤ b ≤ 100?
 */

func main() {
	strMap := make(map[string]int)
	for a := 2; a < 101; a++ {
		for b := 2; b < 101; b++ {
			bigA := big.NewInt(int64(a))
			bigB := big.NewInt(int64(b))
			bigA.Exp(bigA, bigB, nil)
			strMap[bigA.String()] = 0
		}
	}
	println(len(strMap))
}
