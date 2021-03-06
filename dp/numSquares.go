package dp

func numSquares(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = i
	}
	for i := 2; i < n+1; i++ {
		for j := 1; j < i; j++ {
			if i-j*j >= 0 {
				dp[i] = min(dp[i], dp[i-j*j]+1)
			} else {
				break
			}
		}
	}
	return dp[n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
