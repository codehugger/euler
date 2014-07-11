package main

import (
	"fmt"
)

func main() {
	found := true
	for i := 1; i < 1000000000; i++ {
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				found = false
				break
			}

			if j == 20 {
				found = true
			}
		}
		if found == true {
			fmt.Println(i)
			break
		}
	}
}
