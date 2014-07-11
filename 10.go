package main

import (
	"fmt"
	"math"
)

func sieve(n int) []bool {
	values := make([]bool, n)

	for i := range values {
		values[i] = true
	}

	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if values[i] == true {
			for j := i * i; j < n; j += i {
				values[j] = false
			}
		}
	}

	return values
}

func main() {
	sum := 0
	for i, val := range sieve(2000000) {
		if i < 2 {
			continue
		}
		if val == true {
			fmt.Println(i)
			sum += i
		}
	}

	fmt.Println(sum)
}
