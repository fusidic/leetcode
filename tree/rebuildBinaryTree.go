package tree

import (
	"container/list"

	"github.com/fusidic/leetcode/structures"
)

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

// method with stack
func buildTree(preorder []int, inorder []int) *structures.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	stack := list.New()
	root := &structures.TreeNode{}
	root.Val = preorder[0]
	cur := root
	j := 0
	for i := 1; i < len(preorder); i++ {
		// 判断 preorder 中相邻两元素的关系
		if cur.Val != inorder[j] {
			// 与 inorder 中不一致(非父节点)，相邻后点为左子树
			cur.Left = &structures.TreeNode{}
			cur.Left.Val = preorder[i]
			stack.PushBack(cur)
			cur = cur.Left
		} else {
			// 判断该点为父节点，开始回归
			j++
			for stack.Len() != 0 && stack.Back().Value.(*structures.TreeNode).Val == inorder[j] {
				cur = stack.Back().Value.(*structures.TreeNode)
				stack.Remove(stack.Back())
				j++
			}
			cur.Right = &structures.TreeNode{}
			cur.Right.Val = preorder[i]
			cur = cur.Right
		}
	}
	return root
}
