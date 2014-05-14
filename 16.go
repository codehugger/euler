package main

import (
	"fmt"
	"math/big"
	"strconv"
)

func main() {
	fmt.Println("Problem 16")

	prod := big.NewInt(2)
	two := big.NewInt(2)

	for i := 1; i < 1000; i++ {
		prod = prod.Mul(prod, two)
	}

	prodStr := prod.String()

	fmt.Println(prodStr)

	sum := 0

	for _, c := range prodStr {
		num, err := strconv.ParseInt(string(c), 10, 32)
		if err != nil {
			panic(err)
		} else {
			sum += int(num)
		}
	}

	fmt.Println(sum)
}
