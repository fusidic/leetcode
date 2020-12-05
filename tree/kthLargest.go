package tree

import "github.com/fusidic/leetcode/structures"

func kthLargest(root *structures.TreeNode, k int) int {
	var target int
	helper(root, &k, &target)
	return target
}

func helper(root *structures.TreeNode, k *int, target *int) {
	if root.Right != nil {
		helper(root.Right, k, target)
	}

	*k = *k - 1
	if *k == 0 {
		*target = root.Val
		return
	}

	if root.Left != nil {
		helper(root.Left, k, target)
	}
}
