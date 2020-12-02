package tree

import (
	"container/list"

	"github.com/fusidic/leetcode/structures"
)

// https://leetcode-cn.com/problems/cong-shang-dao-xia-da-yin-er-cha-shu-lcof/

func levelOrder(root *structures.TreeNode) []int {
	res := make([]int, 0, 50)
	treeList := list.New()
	treeList.PushBack(root)
	for treeList.Len() != 0 {
		treeList = getNextLevel(treeList, &res)
	}
	return res
}

func getNextLevel(l *list.List, res *[]int) *list.List {
	nextLevel := list.New()
	len := l.Len()
	for i := 0; i < len; i++ {
		treeNode := l.Front().Value.(*structures.TreeNode)
		if treeNode == nil {
			break
		} else {
			*res = append(*res, treeNode.Val)
			if treeNode.Left != nil {
				nextLevel.PushBack(treeNode.Left)
			}
			if treeNode.Right != nil {
				nextLevel.PushBack(treeNode.Right)
			}
			l.Remove(l.Front())
		}

	}
	return nextLevel
}
