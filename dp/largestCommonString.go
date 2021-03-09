package dp

/**
 * longest common substring
 * @param str1 string字符串 the string
 * @param str2 string字符串 the string
 * @return string字符串
 */

// LCS ...
func LCS(str1 string, str2 string) string {
	// write code here
	l1, l2 := len(str1), len(str2)
	max := 0
	idx := 0
	dp := make([][]int, l1+1)
	for i := 0; i < l1+1; i++ {
		dp[i] = make([]int, l2+1)
	}
	for i := 1; i < l1+1; i++ {
		for j := 1; j < l2+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = 0
			}
			if dp[i][j] > max {
				max = dp[i][j]
				idx = i
			}
		}
	}
	if idx > l1-1 {
		return str1[l1-max:]
	} else if max == 0 {
		return "-1"
	}
	return str1[idx-max : idx]
}
