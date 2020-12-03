package tree

import (
	"container/list"

	"github.com/fusidic/leetcode/structures"
)

func levelOrder2(root *structures.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0, 10)
	levelNodeList := list.New()
	levelNodeList.PushBack(root)
	for levelNodeList.Len() != 0 {
		temp := list.New()
		levelValList := make([]int, 0, 10)
		len := levelNodeList.Len()
		for i := 0; i < len; i++ {
			tnode := levelNodeList.Front().Value.(*structures.TreeNode)
			levelValList = append(levelValList, tnode.Val)
			if tnode.Left != nil {
				temp.PushBack(tnode.Left)
			}
			if tnode.Right != nil {
				temp.PushBack(tnode.Right)
			}
			levelNodeList.Remove(levelNodeList.Front())
		}
		levelNodeList = temp
		res = append(res, levelValList)
	}
	return res
}
