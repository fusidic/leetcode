package array

// https://leetcode-cn.com/problems/shun-shi-zhen-da-yin-ju-zhen-lcof/submissions/
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}
	var res []int
	x1, x2 := 0, len(matrix[0])-1
	y1, y2 := 0, len(matrix)-1
	for x1 <= x2 && y1 <= y2 {
		if x1 == x2 && y1 == y2 {
			res = append(res, matrix[y1][x1])
			return res
		}
		for i := x1; i < x2; i++ {
			res = append(res, matrix[y1][i])
		}
		for j := y1; j < y2; j++ {
			res = append(res, matrix[j][x2])
		}

		if y1 != y2 {
			for i := x2; i > x1; i-- {
				res = append(res, matrix[y2][i])
			}
		} else {
			res = append(res, matrix[y2][x2])
		}

		if x1 != x2 {
			for j := y2; j > y1; j-- {
				res = append(res, matrix[j][x1])
			}
		} else {
			res = append(res, matrix[y2][x2])
		}

		x1++
		x2--
		y1++
		y2--
	}
	return res
}
