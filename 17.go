package main

import (
	"fmt"
)

var ones = [20]string{
	"",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
	"eleven",
	"twelve",
	"thirteen",
	"fourteen",
	"fifteen",
	"sixteen",
	"seventeen",
	"eighteen",
	"nineteen",
}

var tens = [10]string{
	"ten",
	"twenty",
	"thirty",
	"forty",
	"fifty",
	"sixty",
	"seventy",
	"eighty",
	"ninety",
}

var hundreds = [10]string{
	"onehundred",
	"twohundred",
	"threehundred",
	"fourhundred",
	"fivehundred",
	"sixhundred",
	"sevenhundred",
	"eighthundred",
	"ninehundred",
}

func numberToExpression(number int) string {
	if number == 1000 {
		return "onethousand"
	}
	if number < 20 {
		return ones[number]
	} else if number < 100 {
		div := number / 10
		rest := number % 10
		return fmt.Sprint(tens[div-1], ones[rest])
	} else {
		div := number / 100
		rest := number % 100
		if rest > 0 {
			return fmt.Sprint(hundreds[div-1], "and", numberToExpression(rest))
		} else {
			return fmt.Sprint(hundreds[div-1])
		}
	}
}

func main() {
	fmt.Println("Problem 17")

	sum := 0

	for i := 1; i <= 1000; i++ {
		sum += len(numberToExpression(i))
		// fmt.Println(i, numberToExpression(i), len(numberToExpression(i)))
	}

	fmt.Println(sum)
}
