package q7

func reverse2(x int) int {
	max := 0x7fffffff
	min := 0x80000000
	var ret int
	for x > 0 {
		ret = ret*10 + x%10
		x /= 10
	}
	if ret > max || ret < min {
		return 0
	}
	return ret
}
