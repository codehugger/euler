package main

import (
	"fmt"
	"github.com/kavehmz/prime"
	"math"
)

var primes []uint64

func isPrime(testNumber uint64) bool {
	// fmt.Printf("Testing: %d\n", testNumber)
	if testNumber < 0 {
		return false
	}
	for i := 0; primes[i] <= testNumber; i++ {
		if primes[i] == testNumber {
			return true
		}
	}
	return false
}

func main() {
	/*
	 * Euler discovered the remarkable quadratic formula: n^2 + n + 41
	 *
	 * It turns out that the formula will produce 40 primes for the
	 * consecutive integer values 0 <= n <= 39. However, when
	 *
	 * n = 40, 40^2 + 40 + 41 = 40(40 + 1) + 41 is divisible by 41,
	 * and certainly when n = 41, 41^2 + 41 + 41 is clearly divisible by
	 * 41.
	 *
	 * The incredible formula n^2 - 79n + 1601 was discovered, which produces
	 * 80 primes for the consecutive values 0 <= n <= 79. The products of the
	 * coefficients, -79 and 1601, is -126479.
	 *
	 * Considering quadratics of the form:
	 *
	 * n^2 + an + b, where |a| < 1000 and |b| <= 1000
	 *
	 * where |n| is the modules/absolute value of n e.g. |11| = 11 and |-4| = 4
	 *
	 * Find the product of the coefficients, a and b, for the quadratic
	 * expression that produces the maximum number of primes for consecutive
	 * values of n, starting with n = 0.
	 */

	// Init primes
	primes = prime.Primes(87400)

	// fmt.Println(primes)

	aMax := 0
	bMax := 0
	nMax := 0

	for a := -1000; a <= 1000; a++ {
		for b := -1000; b <= 1000; b++ {
			n := 0
			// val := n**2 + a*n + b
			// fmt.Printf("n^2 + a*n + b = %d\n", n^2+a*n+b)
			for isPrime(uint64(math.Abs((float64)(n*n + a*n + b)))) {
				n++
			}

			if n > nMax {
				aMax = a
				bMax = b
				nMax = n
			}
		}
	}

	fmt.Printf("A sequence of length %d, is generated by a=%d, b=%d, the product is %d\n",
		nMax, aMax, bMax, aMax*bMax)
}
