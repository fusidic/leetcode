package string

import (
	"strconv"
)

func convertToBin(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
