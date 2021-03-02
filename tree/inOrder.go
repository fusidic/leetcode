package tree

import (
	"container/list"

	"github.com/fusidic/leetcode/structures"
	teststruc "github.com/halfrost/LeetCode-Go/structures"
)

func inorderTraversal(root *structures.TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack := list.New()
	res := make([]int, 0, 20)

	for root != nil || stack.Len() != 0 {
		if root != nil {
			stack.PushBack(root)
			root = root.Left
		} else {
			root = stack.Back().Value.(*structures.TreeNode)
			res = append(res, root.Val)
			root = root.Right
			stack.Remove(stack.Back())
		}
	}
	return res
}

func inorderTraversal2(root *teststruc.TreeNode) []int {
	res := make([]int, 0, 20)
	var inOrder func(root *teststruc.TreeNode)
	inOrder = func(root *teststruc.TreeNode) {
		if root == nil {
			return
		}
		inOrder(root.Left)
		res = append(res, root.Val)
		inOrder(root.Right)
		return
	}

	inOrder(root)
	return res

}
