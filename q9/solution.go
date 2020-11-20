package q9

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	m := 0
	n := x
	for n != 0 {
		m = m*10
		m += n%10
		n = n/10
	}
	if m == x {
		return true
	}
	return false
}