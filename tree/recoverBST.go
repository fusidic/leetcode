package tree

import (
	"math"

	"github.com/fusidic/leetcode/structures"
)

// https://leetcode-cn.com/problems/recover-binary-search-tree/

func recoverTree(root *structures.TreeNode) {
	pre := &structures.TreeNode{Val: math.MinInt64}
	violate := make([]*structures.TreeNode, 0, 4)
	// 获取 BST 中 violation 的情况并存储到数组中
	checkBST(root, &pre, &violate)
	// 由于题目规定只需要交换一次
	// 那么最根本的原因一定是因为 violate 中的末两位逆序
	l := len(violate)
	temp := violate[l-1].Val
	violate[l-1].Val = violate[l-2].Val
	violate[l-2].Val = temp
}

func checkBST(root *structures.TreeNode, pre **structures.TreeNode, violate *[]*structures.TreeNode) {
	if root == nil {
		return
	}
	checkBST(root.Left, pre, violate)
	if root.Val > (*pre).Val {
		*pre = root
		checkBST(root.Right, pre, violate)
	} else {
		*violate = append(*violate, *pre)
		*violate = append(*violate, root)
		checkBST(root.Right, pre, violate)
	}
}
