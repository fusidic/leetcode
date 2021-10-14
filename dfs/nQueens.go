package dfs

func solveNQueens(n int) [][]string {
	res := make([][]string, n)
	for i := 0; i < n; i++ {
		res[i] = make([]string, n)
	}
	var dfs func([]string, int)
	dfs = func(board []string, row int) {
		if row == n {
			t := make([]string, n)
			copy(t, board)
			res = append(res, board)
		}

		for i := 0; i < n; i++ {
			if !validation(board, row, i, n) {
				continue
			}
			board = append(board, string(buildBoard(n, i)))
			dfs(board, row+1)
			board = board[:len(board)-1]
		}
	}
	dfs([]string{}, 0)
	return res
}

func validation(board []string, row, col int, n int) bool {
	// 判断同列
	for i := 0; i < row; i++ {
		if board[i][col] == 'Q' {
			return false
		}
	}

	// 判断45度
	i, j := row-1, col+1
	for i >= 0 && col < n {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j++
	}

	// 判断135度
	i, j = row-1, col-1
	for i >= 0 && j >= 0 {
		if board[i][j] == 'Q' {
			return false
		}
		i--
		j--
	}
	return true
}

func buildBoard(n int, idx int) []byte {
	ret := []byte{}
	for i := 0; i < n; i++ {
		ret = append(ret, '.')
	}
	ret[idx] = 'Q'
	return ret
}
