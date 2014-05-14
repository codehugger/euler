package main

import (
	"fmt"
)

func collatz_chain(n int) []int {
	if n == 1 {
		retval := make([]int, 1)
		retval[0] = 1
		return retval
	} else if n%2 == 0 {
		return append(collatz_chain(n/2), n)
	} else {
		return append(collatz_chain(3*n+1), n)
	}
}

func main() {
	fmt.Println("Problem 14")

	count := 1000000
	max_length := 1
	max_number := 1
	for i := 1; i < count; i++ {
		chain := collatz_chain(i)

		if len(chain) > max_length {
			max_length = len(chain)
			max_number = i
		}
	}

	fmt.Println(max_number)
}
