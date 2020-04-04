package q7

import "strconv"

func reverse(x int) int {
	str := strconv.Itoa(x)
	if str == "-247483648" {
		return 0
	}
	l := len(str)
	var isPositive int
	if string(str[0]) == "-" {
		isPositive = 1
	} else {
		isPositive = -1
		str = str[1:]
	}

}
