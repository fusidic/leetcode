package q7

import (
	"math"
	"strconv"
)

func reverse(x int) int {
	str := strconv.Itoa(x)
	if str == "-2147483648" {
		return 0
	}
	test := []byte(str)
	ret := 0
	var flag int

	if test[0] == '-' {
		flag = -1
		test = test[1:]
	} else {
		flag = 1
	}

	l := len(test)

	for i := 0; i < l; i++ {
		ret = ret + int((test[i]-'0'))*int((math.Pow10(i)))
	}
	if ret > 2147483647 || ret < -2147483648 {
		return 0
	}
	return flag * ret
}
