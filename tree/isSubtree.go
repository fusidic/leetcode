package tree

import "github.com/fusidic/leetcode/structures"

// https://leetcode-cn.com/problems/shu-de-zi-jie-gou-lcof/

func isSubStructure(A *structures.TreeNode, B *structures.TreeNode) bool {
	if B == nil || A == nil {
		return false
	}
	return subTreeCheck(A, B) || isSubStructure(A.Left, B) || isSubStructure(A.Right, B)
}

func subTreeCheck(A, B *structures.TreeNode) bool {
	if B == nil {
		return true
	}
	if A == nil {
		return false
	}
	if A.Val == B.Val {
		return subTreeCheck(A.Left, B.Left) && subTreeCheck(A.Right, B.Right)
	}
	return false
}

// 另一种思路：直接将数转换为 list, 直接将判断子树的问题转换为子字符串的问题。
