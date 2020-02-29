package main

import (
	"fmt"
)

func main() {
	fmt.Println(factorial(3))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
