package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := -63434
	s := strconv.Itoa(x)
	fmt.Println(s)
	for _, c := range s {
		fmt.Println(c)
	}
}

func reverse(x int) int {
	return 0
}
