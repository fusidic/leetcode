package tree

import (
	"container/list"

	"github.com/fusidic/leetcode/structures"
)

func zigzagLevelOrder(root *structures.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	level := 0
	nodeList := list.New()
	nodeList.PushBack(root)
	res := make([][]int, 0, 5)
	for nodeList.Len() != 0 {
		len := nodeList.Len()
		temp := make([]int, 0, 10)
		nextLevel := list.New()

		// level the travel
		for i := 0; i < len; i++ {
			tnode := nodeList.Front().Value.(*structures.TreeNode)
			temp = append(temp, tnode.Val)
			if level%2 == 0 {
				// 逆序填入
				if tnode.Left != nil {
					nextLevel.PushFront(tnode.Left)
				}
				if tnode.Right != nil {
					nextLevel.PushFront(tnode.Right)
				}

			} else {
				if tnode.Right != nil {
					nextLevel.PushFront(tnode.Right)
				}

				if tnode.Left != nil {
					nextLevel.PushFront(tnode.Left)
				}
			}
			nodeList.Remove(nodeList.Front())
		}
		res = append(res, temp)
		nodeList = nextLevel
		level++
	}
	return res
}
