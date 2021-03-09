package dp

func minSteps(n int) int {
	if n == 1 {
		return 0
	}
	dp := make([]int, n+1)
	for i := 2; i < n+1; i++ {
		for k := i / 2; k > 0; k-- {
			if i%k == 0 {
				dp[i] = dp[k] + i/k
				break
				// k = i / 2; k > 0; k--){
				// if(i % k == 0){
				//     dp[i] = dp[k] + 1 + (i - k) / k;
				//     break;
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
