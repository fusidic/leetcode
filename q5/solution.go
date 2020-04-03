func longestPalindrome(s string) string {
	i := 0
	l := ""
	temp := ""
	for i < len(s) {
		temp = getPalindrome(s, i, i)
		if len(l) < len(temp) {
			l = temp
		}
		if i+1 < len(s) {
			temp = getPalindrome(s, i, i+1)
			if len(l) < len(temp) {
				l = temp
			}
		}
		i++
	}
	return l
}

func getPalindrome(s string, left int, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			left++
			right--
			break
		}
	}
	if left < 0 || right >= len(s) {
		return s[left+1 : right]
	}
	return s[left:right+1]
}