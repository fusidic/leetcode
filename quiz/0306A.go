package quiz

import "fmt"

func numOfWays(n int) int {
	if n == 1 {
		return 12
	}
	typeA := 6
	typeB := 6
	for i := 1; i < n; i++ {
		A := (typeA*3 + typeB*2) % (1e9 + 7)
		B := (typeA*2 + typeB*2) % (1e9 + 7)
		typeA = A
		typeB = B
	}
	return (typeA + typeB) % (1e9 + 7)
}

func test() {
	var n []int
	fmt.Scan(&n)
	for i := 0; i < len(n); i++ {
		numOfWays(n[i])
	}
}
