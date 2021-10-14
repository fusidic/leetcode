package netease

func papers(ages []int) int {
	papers := make([]int, len(ages))
	sum := 0
	papers[0] = 1
	// 从左到右贪心
	for i := 1; i < len(ages); i++ {
		if ages[i] > ages[i-1] {
			papers[i] = papers[i-1] + 1
		} else {
			papers[i] = 1
		}
	}
	// 从右到左贪心
	for i := len(ages) - 2; i >= 0; i-- {
		if ages[i] > ages[i+1] {
			papers[i] = max(papers[i], papers[i+1]+1)
		}
	}
	for _, v := range papers {
		sum += v
	}
	return sum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
