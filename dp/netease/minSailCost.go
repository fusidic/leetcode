package netease

func minSailCost(input [][]int) int {
	m := len(input)
	if m == 0 {
		return 0
	}
	n := len(input[0])
	dp := make([][]int, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]int, n+1)
	}

	// road
	roadCost := make(map[int]int)
	roadCost[0] = 2
	roadCost[1] = 1

	// 初始化边界
	for i := 0; i < m+1; i++ {
		dp[i][0] = 1 << 32
	}
	for j := 0; j < n+1; j++ {
		dp[0][j] = 1 << 32
	}
	dp[1][1] = 0

	// dp
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			if i == 1 && j == 1 {
				continue
			}
			if input[i-1][j-1] != 2 {
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + roadCost[input[i-1][j-1]]
				continue
			}
			dp[i][j] = 1 << 32
		}
	}
	if dp[m][n] >= 1<<32 {
		return -1
	}
	return dp[m][n]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
