package tree

import "github.com/halfrost/LeetCode-Go/structures"

func lowestCommonAncestor(root, p, q *structures.TreeNode) *structures.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

var ancestor *structures.TreeNode

func lowestCommonAncestor2(root, p, q *structures.TreeNode) *structures.TreeNode {
	if root == nil {
		return nil
	}
	jugg(root, p, q)
	return ancestor
}

func jugg(root, p, q *structures.TreeNode) bool {
	if root == nil {
		return false
	}
	stat := false
	if root.Val == p.Val || root.Val == q.Val {
		stat = true
	}
	l := jugg(root.Left, p, q)
	r := jugg(root.Right, p, q)
	if l && r {
		ancestor = root
		return true
	}
	if l || r {
		return true
	}
	if root.Val == p.Val || root.Val == q.Val {
		return true
	}
	return stat
}
