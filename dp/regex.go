package dp

func isMatch(s string, p string) bool {
	x := len(s) + 1
	y := len(p) + 1
	dp := make([][]bool, x)
	for i := 0; i < x; i++ {
		dp[i] = make([]bool, y)
	}
	dp[0][0] = true
	for j := 2; j < y; j += 2 {
		dp[0][j] = dp[0][j-2] && p[j-1] == '*'
	}
	for i := 1; i < x; i++ {
		for j := 1; j < y; j++ {
			if p[j-1] == '*' {
				// 不会越界，首位不可能为 *
				if dp[i][j-2] {
					dp[i][j] = true
				} else if dp[i-1][j] && s[i-1] == p[j-2] {
					dp[i][j] = true
				} else if dp[i-1][j] && p[j-2] == '.' {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			} else {
				if dp[i-1][j-1] && p[j-1] == '.' {
					dp[i][j] = true
				} else if dp[i-1][j-1] && p[j-1] == s[i-1] {
					dp[i][j] = true
				} else {
					dp[i][j] = false
				}
			}
		}
	}

	return dp[x-1][y-1]
}
