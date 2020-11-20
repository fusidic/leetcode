package leetcode

import "strings"

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	rows := make([]strings.Builder, numRows)
	i := 0
	flag := -1
	for _, c := range s {
		rows[i].WriteRune(c)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}
	var res strings.Builder
	for _, row := range rows {
		res.WriteString(row.String())
	}
	return res.String()
}
