package leetcode

import "github.com/fusidic/leetcode/leetutils"

func longestPalindrome2(s string) string {
	reverse := leetutils.Reverse(s)
	maxLen := 0
	index := 0
	l := len(s)
	mtx := make([][]int, l)
	for i := 0; i < l; i++ {
		// 初始化，默认值都为0
		mtx[i] = make([]int, l)
		for j := 0; j < l; j++ {
			// 字符相同
			if s[i] == reverse[j] {
				// 边界处理
				if i == 0 || j == 0 {
					mtx[i][j] = 1
				} else {
					mtx[i][j] = mtx[i-1][j-1] + 1
				}
			}

			// 更新maxLen
			if mtx[i][j] > maxLen {
				// 判断末尾元素在s和reverse中是否是同一个元素
				if (l - 1 - j + mtx[i][j] - 1) == i {
					maxLen = mtx[i][j]
					index = i - maxLen + 1
				}
			}
		}
	}
	return s[index : index+maxLen]
}
