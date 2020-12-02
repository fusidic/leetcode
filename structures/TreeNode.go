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
		temp := getNextLevelList(subtreeList)
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

func getNextLevelList(treeList []*TreeNode) []*TreeNode {
	if treeList == nil {
		return nil
	}
	listLen := len(treeList)
	var subtreeList []*TreeNode
	for i := 0; i < listLen; i++ {
		if treeList[i] != nil {
			if treeList[i].Val == NULL {
				continue
			}
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

// LeetArray2Tree transfer tree array to TreeNode
func LeetArray2Tree(l []int) *TreeNode {
	if len(l) == 0 {
		return nil
	} else if len(l)%2 == 0 {
		// complement tail NULL if exists.
		l = append(l, NULL)
	}
	treeList := []*TreeNode{}
	// Found a rule bewteen PARENTS node and CHILD.
	// During each cycle, it will found the CHILD nodes for index.
	for i := 0; i < len(l); i++ {
		if l[i] != NULL {
			treeList = append(treeList, &TreeNode{
				Val:   l[i],
				Left:  nil,
				Right: nil,
			})
		} else {
			// When index is nil, apparently it's the end edge.
			treeList = append(treeList, nil)
		}

	}
	index := 0
	for j := 1; j < len(l); {
		if l[index] != NULL {
			treeList[index].Left = treeList[j]
			treeList[index].Right = treeList[j+1]
			index++
			j += 2
		} else {
			index++
		}
	}
	return treeList[0]
}
