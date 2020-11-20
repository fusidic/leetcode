package leetcode

import (
	"math"
	"strings"
)

// This is a really stupid method
func myAtoi(str string) int {
	max := 0x7fffffff
	min := -0x80000000
	var ret int
	i := 0
	l := 0
	flag := 0
	str = strings.TrimSpace(str)
	if str == "" {
		return 0
	}
	if str[0] != '-' && str[0] != '+' && (str[0] < '0' || str[0] > '9') {
		return 0
	}
	for i = 0; i < len(str); i++ {
		if str[i] == '-' || str[i] == '+' || (str[i] >= '0' && str[i] <= '9') {
			break
		}
	}
	if str[i] == '-' {
		flag = -1
		i++
	} else if str[i] == '+' {
		if len(str) == 1 {
			return 0
		}
		i++
		flag = 1
	} else {
		flag = 1
	}

	start := i
	for i < len(str) {
		if str[i] >= '0' && str[i] <= '9' {
			l++
			i++
		} else {
			break
		}
	}

	for l != 0 {
		tmp := float64(str[start]-'0') * math.Pow10(l-1)
		if tmp*float64(flag) > float64(max) {
			return max
		} else if tmp*float64(flag) < float64(min) {
			return min
		}
		ret += int(tmp)
		l--
		start++
	}

	ret = ret * flag
	if ret > max {
		return max
	} else if ret < min {
		return min
	}
	return ret
}
