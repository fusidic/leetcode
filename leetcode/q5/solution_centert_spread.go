package leetcode

func longestPalindrome1(s string) string {
	i := 0
	l := ""
	temp := ""
	for i < len(s) {
		// 奇数
		temp = getPalindrome(s, i, i)
		if len(l) < len(temp) {
			l = temp
		}
		// 判断是否会越界
		if i+1 < len(s) {
			// 偶数
			temp = getPalindrome(s, i, i+1)
			if len(l) < len(temp) {
				l = temp
			}
		}
		i++
	}
	return l
}

// getPalindrome 向两边延伸找到最大回文串
func getPalindrome(s string, left int, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			left--
			right++
		} else {
			// 左右不相同，回退
			left++
			right--
			break
		}
	}
	// 碰到字符串边界的处理
	if left < 0 || right >= len(s) {
		return s[left+1 : right]
	}
	return s[left : right+1]
}
