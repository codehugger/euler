package main

import (
	"fmt"
	"math"
)

func main() {
	for a := 3; a < 1000; a++ {
		for b := 4; b < 1000; b++ {
			c := int(math.Sqrt(float64(a*a + b*b)))
			sum := a + b + c
			if sum == 1000 && a < b && b < c && (a*a+b*b) == c*c {
				fmt.Printf("%d + %d + %d = %d\n", a, b, c, sum)
				fmt.Println("Product is:", 200*375*425)
			}
		}
	}
}
