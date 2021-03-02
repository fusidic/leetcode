package string

import (
	"strconv"
)

func convertToBin(n int, k int) string {
	str := strconv.FormatInt(int64(n), 2)
	if str[0] == '-' {
		neg := []byte(str)
		neg[0] = '0'
		// 按位取反
		var i int
		for i = 1; i < len(neg); i++ {
			if neg[i] == '0' {
				neg[i] = '1'
			} else {
				neg[i] = '0'
			}
		}
		flag := 1
		i--
		for i >= 0 {
			if flag == 0 {
				break
			}
			if neg[i] == '1' {
				neg[i] = '0'
				flag = 1
			} else {
				neg[i] = '1'
				flag = 0
			}
			i--
		}
		return "1" + genZero(k-len(neg)-1) + string(neg)
	}
	return genZero(k-len(str)) + str
}

func genZero(k int) string {
	str := ""
	for i := 0; i < k; i++ {
		str += "0"
	}
	return str
}
