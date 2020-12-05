package tree

import (
	"github.com/fusidic/leetcode/structures"
)

func isBalanced(root *structures.TreeNode) bool {
	if root == nil {
		return true
	}
	left := getDepth(root.Left)
	right := getDepth(root.Right)
	return dif(left, right) < 2 && isBalanced(root.Left) && isBalanced(root.Right)
}

func getDepth(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + max(getDepth(root.Left), getDepth(root.Right))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func dif(x, y int) int {
	return max(x, y)*2 - x - y
}
