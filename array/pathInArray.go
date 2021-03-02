package array

// Exist ...
func Exist(board [][]byte, word string) bool {
	height := len(board)
	width := len(board[0])
	length := len(word)
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		res := false
		// dfs超出边界
		if i < 0 || i >= height || j < 0 || j >= width {
			return false
		}
		if board[i][j] == word[index] {
			temp := board[i][j]
			board[i][j] = ' '
			if index == length-1 {
				return true
			}
			res = dfs(i-1, j, index+1) || dfs(i+1, j, index+1) || dfs(i, j-1, index+1) || dfs(i, j+1, index+1)
			// 还原现场
			board[i][j] = temp
			return res
		}
		return false
	}

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] == word[0] {
				if dfs(i, j, 0) {
					return true
				}
			}
		}
	}
	return false
}
