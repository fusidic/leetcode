package tree

import (
	"container/list"

	"github.com/fusidic/leetcode/structures"
)

// https://leetcode-cn.com/problems/zhong-jian-er-cha-shu-lcof/

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
		// 相邻为左子树
		if cur.Val != inorder[j] {
			cur.Left = &structures.TreeNode{}
			cur.Left.Val = preorder[i]
			stack.PushBack(cur)
			cur = cur.Left
		} else {
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
