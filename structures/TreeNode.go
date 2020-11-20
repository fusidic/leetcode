package structures

// TreeNode is a node in a binary tree data structure.
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// NULL what can I do ...
var NULL = -1 << 63

// Tree2LeetArray transfer TreeNode to dumb array in LeetCode.
func Tree2LeetArray(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var treeList []*TreeNode
	subtreeList := []*TreeNode{root}
	for subtreeList != nil {
		treeList = append(treeList, subtreeList...)
		temp := getSubtreeList(subtreeList)
		subtreeList = temp
	}

	res := make([]int, len(treeList))
	for i := 0; i < len(treeList); i++ {
		if treeList[i] != nil {
			res[i] = treeList[i].Val
		} else {
			res[i] = NULL
		}
	}
	res = ridTailNull(res)
	return res
}

func getSubtreeList(treeList []*TreeNode) []*TreeNode {
	if treeList == nil {
		return nil
	}
	listLen := len(treeList)
	var subtreeList []*TreeNode
	for i := 0; i < listLen; i++ {
		if treeList[i] != nil {
			subtreeList = append(subtreeList, treeList[i].Left)
			subtreeList = append(subtreeList, treeList[i].Right)
		}
	}
	return subtreeList
}

func ridTailNull(list []int) []int {
	i := len(list) - 1
	for ; i > 0; i-- {
		if list[i] != NULL {
			break
		}
	}
	return list[0 : i+1]
}

// Tree2PreOrder transfer TreeNode to PreOrder slice.
func Tree2PreOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	if root.Left == nil && root.Right == nil {
		return []int{root.Val}
	}

	res := []int{root.Val}
	res = append(res, Tree2PreOrder(root.Left)...)
	res = append(res, Tree2PreOrder(root.Right)...)
	return res
}
