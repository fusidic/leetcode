package array

func combinationSum(candidates []int, target int) [][]int {
	list := []int{}
	res := [][]int{}
	var dfs func(target, idx int)
	dfs = func(target, idx int) {
		if idx == len(candidates) {
			return
		}

		if target == 0 {
			// append
			// res = append(res, append([]int{}, list...))
			res = append(res, list)
			return
		}
		// 不选
		dfs(target, idx+1)
		// 选
		if target-candidates[idx] < 0 {
			return
		}
		if target-candidates[idx] >= 0 {
			list = append(list, candidates[idx])
			dfs(target-candidates[idx], idx)
			// 回溯
			list = list[:len(list)-1]
		}

	}
	dfs(target, 0)
	return res
}
