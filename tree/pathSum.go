package tree

import (
	"github.com/fusidic/leetcode/structures"
)

// https://leetcode-cn.com/problems/er-cha-shu-zhong-he-wei-mou-yi-zhi-de-lu-jing-lcof/comments/
func pathSum(root *structures.TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0)
	// pre-allocating saves time.
	valList := make([]int, 0, 8)
	getPathSum(root, sum, valList, &res)
	return res
}

func getPathSum(root *structures.TreeNode, pathSum int, valList []int, res *[][]int) {
	if root == nil {
		return
	}
	valList = append(valList, root.Val)
	pathSum -= root.Val
	if pathSum == 0 && root.Left == nil && root.Right == nil {
		tmp := make([]int, len(valList))
		copy(tmp, valList)
		*res = append(*res, tmp)
	} else {
		getPathSum(root.Left, pathSum, valList, res)
		getPathSum(root.Right, pathSum, valList, res)
	}
	// Yes, it really did.
	// fmt.Printf("valList: %p %d \t res: %p %d\n", valList, valList, res, res)
	valList = valList[:len(valList)-1]
}
