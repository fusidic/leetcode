package tree

import (
	"github.com/fusidic/leetcode/structures"
)

// https://leetcode-cn.com/problems/unique-binary-search-trees-ii/
func generateTrees(n int) []*structures.TreeNode {
	if n == 0 {
		return []*structures.TreeNode{}
	}
	return generateBST(1, n)
}

func generateBST(x, y int) []*structures.TreeNode {

	res := make([]*structures.TreeNode, 0, 10)
	if x > y {
		res = append(res, nil)
	}
	if x == y {
		return []*structures.TreeNode{{Val: x}}
	}
	for i := x; i <= y; i++ {
		left := generateBST(x, i-1)
		right := generateBST(i+1, y)
		for _, l := range left {
			for _, r := range right {
				root := &structures.TreeNode{Val: i}
				root.Left = l
				root.Right = r
				res = append(res, root)
			}
		}
	}
	return res
}
